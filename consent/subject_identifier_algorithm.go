// Copyright © 2022 Ory Corp

package consent

import "github.com/ory/hydra/client"

type SubjectIdentifierAlgorithm interface {
	// Obfuscate derives a pairwise subject identifier from the given string.
	Obfuscate(subject string, client *client.Client) (string, error)
}
