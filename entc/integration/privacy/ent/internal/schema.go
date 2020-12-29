// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/facebook/ent/entc/integration/privacy/ent/schema","Package":"github.com/facebook/ent/entc/integration/privacy/ent","Schemas":[{"name":"Task","config":{"Table":""},"edges":[{"name":"teams","type":"Team"},{"name":"owner","type":"User","ref_name":"tasks","unique":true,"inverse":true}],"fields":[{"name":"title","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"status","type":{"Type":6,"Ident":"task.Status","PkgPath":"","Nillable":false,"RType":null},"enums":[{"N":"planned","V":"planned"},{"N":"in_progress","V":"in_progress"},{"N":"closed","V":"closed"}],"default":true,"default_value":"planned","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"uuid","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","Nillable":true,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"hooks":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":true,"MixinIndex":1},{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"Team","config":{"Table":""},"edges":[{"name":"tasks","type":"Task","ref_name":"teams","inverse":true},{"name":"users","type":"User","ref_name":"teams","inverse":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"User","config":{"Table":""},"edges":[{"name":"teams","type":"Team"},{"name":"tasks","type":"Task"}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"unique":true,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"age","type":{"Type":17,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":true,"MixinIndex":1},{"Index":0,"MixedIn":false,"MixinIndex":0}]}],"Features":["privacy","entql","schema/snapshot"]}`
