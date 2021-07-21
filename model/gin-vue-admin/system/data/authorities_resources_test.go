package model

import "testing"

func TestAuthorityResources_TableName(t *testing.T) {
	var entity AuthorityResources
	t.Log(entity.TableName())
}