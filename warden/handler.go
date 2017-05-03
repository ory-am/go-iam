package warden

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/hydra/config"
	"github.com/ory/hydra/firewall"
	"github.com/pkg/errors"
)

const (
	// TokenAllowedHandlerPath points to the token access request validation endpoint.
	TokenAllowedHandlerPath = "/warden/token/allowed"

	// AllowedHandlerPath points to the access request validation endpoint.
	AllowedHandlerPath = "/warden/allowed"
)

type wardenAuthorizedRequest struct {
	Scopes []string `json:"scopes"`
	Token  string   `json:"token"`
}

type wardenAccessRequest struct {
	*firewall.TokenAccessRequest
	*wardenAuthorizedRequest
}

var notAllowed = struct {
	Allowed bool `json:"allowed"`
}{Allowed: false}

var invalid = struct {
	Valid bool `json:"valid"`
}{Valid: false}

// WardenHandler is capable of handling HTTP request and validating access tokens and access requests.
type WardenHandler struct {
	H      herodot.Writer
	Warden firewall.Firewall
}

func NewHandler(c *config.Config, router *httprouter.Router) *WardenHandler {
	ctx := c.Context()

	h := &WardenHandler{
		H:      herodot.NewJSONWriter(c.GetLogger()),
		Warden: ctx.Warden,
	}
	h.SetRoutes(router)

	return h
}

func (h *WardenHandler) SetRoutes(r *httprouter.Router) {
	r.POST(TokenAllowedHandlerPath, h.TokenAllowed)
	r.POST(AllowedHandlerPath, h.Allowed)
}

func (h *WardenHandler) Allowed(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var ctx = r.Context()
	if _, err := h.Warden.TokenAllowed(ctx, h.Warden.TokenFromRequest(r), &firewall.TokenAccessRequest{
		Resource: "rn:hydra:warden:allowed",
		Action:   "decide",
	}, "hydra.warden"); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	var access = new(firewall.AccessRequest)
	if err := json.NewDecoder(r.Body).Decode(access); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}
	defer r.Body.Close()

	if err := h.Warden.IsAllowed(ctx, access); err != nil {
		h.H.Write(w, r, &notAllowed)
		return
	}

	res := notAllowed
	res.Allowed = true
	h.H.Write(w, r, &res)
}

func (h *WardenHandler) TokenAllowed(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	_, err := h.Warden.TokenAllowed(ctx, h.Warden.TokenFromRequest(r), &firewall.TokenAccessRequest{
		Resource: "rn:hydra:warden:token:allowed",
		Action:   "decide",
	}, "hydra.warden")
	if err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	var ar = wardenAccessRequest{
		TokenAccessRequest:      new(firewall.TokenAccessRequest),
		wardenAuthorizedRequest: new(wardenAuthorizedRequest),
	}
	if err := json.NewDecoder(r.Body).Decode(&ar); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}
	defer r.Body.Close()

	authContext, err := h.Warden.TokenAllowed(ctx, ar.Token, ar.TokenAccessRequest, ar.Scopes...)
	if err != nil {
		h.H.Write(w, r, &notAllowed)
		return
	}

	h.H.Write(w, r, struct {
		*firewall.Context
		Allowed bool `json:"allowed"`
	}{
		Context: authContext,
		Allowed: true,
	})
}

func TokenFromRequest(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	split := strings.SplitN(auth, " ", 2)
	if len(split) != 2 || !strings.EqualFold(split[0], "bearer") {
		return ""
	}

	return split[1]
}
