cd /workspaces/SE-63/backend
go run main.go

cd /workspaces/SE-63/frontend
yarn install
# yarn start
NODE_OPTIONS=--openssl-legacy-provider yarn workspace app start
