run-registrar-test: 
	cd ./examples/vdr && go run .

run-jwt-test: 
	@go run ./examples/vc/jwt

run-scenario: 
	@go run ./actors/holder/cmd/main.go

# Registrar port: 9000
run-registrar: 
	@go run vdr/registrar/registrar.go

# Resover port: 9001
run-resolver: 
	@go run vdr/resolver/resolver.go

# Issuer ROT port: 1120
run-root-issuer: 
	@go run actors/issuer/RootOfTrustIssuer/cmd/main.go

# Issuer University port: 1121
run-university-issuer: 
	@go run actors/issuer/UniversityIssuer/cmd/main.go
