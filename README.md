# Go-REST
Example of How To Build RESTful Services in Golang

## Starting The Server

> go run restapi.go

## Adding A New Review

> curl -X POST -d '{ "title" : "$title", "author" : "$author", "gitdiff" : "$gitdiff" }' localhost:6960/api/reviews/new

## Querying Review Data

> curl localhost:6960/api/reviews/$review_id

## Updating An Existing Review

> curl -X POST -d '{ "gitdiff" : "$gitdiff" }' localhost:6960/api/reviews/$review_id/update

## Additional Notes

This project includes a "vimrc.example" file for developing with Golang in vim. It provides code formatting, syntax highlighting and auto-importing of packages and requires the following projects: 

- https://github.com/fatih/vim-go
- https://github.com/VundleVim/Vundle.vim

Once installed, rename the "vimrc.example" file to ".vimrc" and place it in your home directory.
