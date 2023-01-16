# changes in OpenAPI config

- modify `one of` response array to an array of a unified object to enable the generator to create a proper struct for the response
- run: `swagger-cli bundle -r -o costs.json swagger.json`
- if the array was modified to object by the swagger-cli -> modify it back to an array
