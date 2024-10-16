# Booksties

docker build -f src/Services/User/Identity.API/Dockerfile -t test1 .

docker build -f src/Services/Catalog/CatalogService/Dockerfile -t test1 .

docker run -d -p 8080:8080 --name user-api-container test1

docker-compose -f "docker-compose.yml" up -d

docker-compose up --build

// "user-route": {
// "ClusterId": "user-cluster",
// "AuthorizationPolicy": "default",
// "Match": {
// "Path": "/user-service/{**catch-all}",
// "Methods": [ "GET", "POST" ]
// },
// "Transforms": [
// {
// "PathPattern": "{**catch-all}"
// }
// ]
// },
// "auth": {
// "ClusterId": "user-cluster",
// "Match": {
// "Path": "/{action}/{**catch-all}",
// "Methods": [ "GET", "POST" ]
// },
// "Transforms": [
// {
// "PathPattern": "{action}/{**catch-all}"
// }
// ]
// },
