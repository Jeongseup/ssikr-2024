# ssikr-2024

시나리오: 사용자인 holder는 RootOfTrust으로 부터 대면인증을 통해, ID VC를 발급받은 후 University로 가서 재학증명서를 발급받는다.

```bash
# setup
go install github.com/ddollar/forego@latest
# run registrar, resolver, issuers
forego start

# run scenario
make run-scenario
```

### Reference

- https://github.com/SSI-Korea/book-examples
