# Changes

1. error generating code: error creating operation definitions: path '/v1/projects/{projectID}/instances/{instanceId}/credentials/{credentialsID}' refers to parameter 'instanceId', which doesn't exist in specification -> fix by removing ref and copying over the actual schema of instanceId
