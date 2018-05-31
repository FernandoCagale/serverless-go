# serverless-go
serverless-go


- [x] Build serverless
- [ ] Stage serverless
- [ ] Environment serverless
- [ ] PostgreSQL(GORM)
- [ ] JWT
- [ ] Test
- [ ] CI(Gitlab)


`Install serverless globally`

```sh
$ npm install serverless -g
```

`Add credentials`

```sh
$ serverless config credentials --provider aws --key AWS_KEY --secret AWS_SECRET
```

`Add dependencies`

```sh
$ npm install
```

`Install dep`

```sh
$ go get -v github.com/golang/dep/cmd/dep
```

`Install the project's dependencies`

```sh
$ dep ensure
```

`Deploy`

```sh
$ serverless deploy
```