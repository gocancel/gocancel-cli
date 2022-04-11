#!/bin/bash

# generate mocks from wrapped services

set -euo pipefail

cd "client"

go get -u github.com/golang/mock/mockgen

mockgen -source categories.go -package=mocks CategoriesService > mocks/CategoriesService.go
mockgen -source letters.go -package=mocks LettersService > mocks/LettersService.go
mockgen -source organizations.go -package=mocks OrganizationsService > mocks/OrganizationsService.go
mockgen -source products.go -package=mocks ProductsService > mocks/ProductsService.go
