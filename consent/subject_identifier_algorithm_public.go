// Copyright © 2022 Ory Corp

package consent

import "github.com/ory/hydra/client"

type SubjectIdentifierAlgorithmPublic struct{}

func NewSubjectIdentifierAlgorithmPublic() *SubjectIdentifierAlgorithmPublic {
	return &SubjectIdentifierAlgorithmPublic{}
}

func (g *SubjectIdentifierAlgorithmPublic) Obfuscate(subject string, client *client.Client) (string, error) {
	return subject, nil
}
