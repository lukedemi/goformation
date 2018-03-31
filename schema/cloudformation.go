package schema

// cloudformationSchema defined a JSON Schema that can be used to validate CloudFormation/SAM templates
var cloudformationSchema = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "additionalProperties": false,
    "definitions": {
        "AWS::ApiGateway::Account": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CloudWatchRoleArn": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Account"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::ApiKey": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CustomerId": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Enabled": {
                            "type": "boolean"
                        },
                        "GenerateDistinctId": {
                            "type": "boolean"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "StageKeys": {
                            "items": {
                                "$ref": "#/definitions/AWS::ApiGateway::ApiKey.StageKey"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::ApiKey"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::ApiKey.StageKey": {
            "additionalProperties": false,
            "properties": {
                "RestApiId": {
                    "type": "string"
                },
                "StageName": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::Authorizer": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AuthType": {
                            "type": "string"
                        },
                        "AuthorizerCredentials": {
                            "type": "string"
                        },
                        "AuthorizerResultTtlInSeconds": {
                            "type": "number"
                        },
                        "AuthorizerUri": {
                            "type": "string"
                        },
                        "IdentitySource": {
                            "type": "string"
                        },
                        "IdentityValidationExpression": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "ProviderARNs": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Authorizer"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::BasePathMapping": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "BasePath": {
                            "type": "string"
                        },
                        "DomainName": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "Stage": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DomainName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::BasePathMapping"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::ClientCertificate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::ClientCertificate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Deployment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "StageDescription": {
                            "$ref": "#/definitions/AWS::ApiGateway::Deployment.StageDescription"
                        },
                        "StageName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Deployment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Deployment.MethodSetting": {
            "additionalProperties": false,
            "properties": {
                "CacheDataEncrypted": {
                    "type": "boolean"
                },
                "CacheTtlInSeconds": {
                    "type": "number"
                },
                "CachingEnabled": {
                    "type": "boolean"
                },
                "DataTraceEnabled": {
                    "type": "boolean"
                },
                "HttpMethod": {
                    "type": "string"
                },
                "LoggingLevel": {
                    "type": "string"
                },
                "MetricsEnabled": {
                    "type": "boolean"
                },
                "ResourcePath": {
                    "type": "string"
                },
                "ThrottlingBurstLimit": {
                    "type": "number"
                },
                "ThrottlingRateLimit": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::Deployment.StageDescription": {
            "additionalProperties": false,
            "properties": {
                "CacheClusterEnabled": {
                    "type": "boolean"
                },
                "CacheClusterSize": {
                    "type": "string"
                },
                "CacheDataEncrypted": {
                    "type": "boolean"
                },
                "CacheTtlInSeconds": {
                    "type": "number"
                },
                "CachingEnabled": {
                    "type": "boolean"
                },
                "ClientCertificateId": {
                    "type": "string"
                },
                "DataTraceEnabled": {
                    "type": "boolean"
                },
                "Description": {
                    "type": "string"
                },
                "DocumentationVersion": {
                    "type": "string"
                },
                "LoggingLevel": {
                    "type": "string"
                },
                "MethodSettings": {
                    "items": {
                        "$ref": "#/definitions/AWS::ApiGateway::Deployment.MethodSetting"
                    },
                    "type": "array"
                },
                "MetricsEnabled": {
                    "type": "boolean"
                },
                "ThrottlingBurstLimit": {
                    "type": "number"
                },
                "ThrottlingRateLimit": {
                    "type": "number"
                },
                "Variables": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::DocumentationPart": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Location": {
                            "$ref": "#/definitions/AWS::ApiGateway::DocumentationPart.Location"
                        },
                        "Properties": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Location",
                        "Properties",
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::DocumentationPart"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::DocumentationPart.Location": {
            "additionalProperties": false,
            "properties": {
                "Method": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Path": {
                    "type": "string"
                },
                "StatusCode": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::DocumentationVersion": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "DocumentationVersion": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DocumentationVersion",
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::DocumentationVersion"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::DomainName": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CertificateArn": {
                            "type": "string"
                        },
                        "DomainName": {
                            "type": "string"
                        },
                        "EndpointConfiguration": {
                            "$ref": "#/definitions/AWS::ApiGateway::DomainName.EndpointConfiguration"
                        },
                        "RegionalCertificateArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DomainName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::DomainName"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::DomainName.EndpointConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Types": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::GatewayResponse": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ResponseParameters": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "ResponseTemplates": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "ResponseType": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "StatusCode": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ResponseType",
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::GatewayResponse"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Method": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApiKeyRequired": {
                            "type": "boolean"
                        },
                        "AuthorizationType": {
                            "type": "string"
                        },
                        "AuthorizerId": {
                            "type": "string"
                        },
                        "HttpMethod": {
                            "type": "string"
                        },
                        "Integration": {
                            "$ref": "#/definitions/AWS::ApiGateway::Method.Integration"
                        },
                        "MethodResponses": {
                            "items": {
                                "$ref": "#/definitions/AWS::ApiGateway::Method.MethodResponse"
                            },
                            "type": "array"
                        },
                        "OperationName": {
                            "type": "string"
                        },
                        "RequestModels": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "RequestParameters": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "boolean"
                                }
                            },
                            "type": "object"
                        },
                        "RequestValidatorId": {
                            "type": "string"
                        },
                        "ResourceId": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "HttpMethod",
                        "ResourceId",
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Method"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Method.Integration": {
            "additionalProperties": false,
            "properties": {
                "CacheKeyParameters": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "CacheNamespace": {
                    "type": "string"
                },
                "ContentHandling": {
                    "type": "string"
                },
                "Credentials": {
                    "type": "string"
                },
                "IntegrationHttpMethod": {
                    "type": "string"
                },
                "IntegrationResponses": {
                    "items": {
                        "$ref": "#/definitions/AWS::ApiGateway::Method.IntegrationResponse"
                    },
                    "type": "array"
                },
                "PassthroughBehavior": {
                    "type": "string"
                },
                "RequestParameters": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "RequestTemplates": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "type": "string"
                },
                "Uri": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::Method.IntegrationResponse": {
            "additionalProperties": false,
            "properties": {
                "ContentHandling": {
                    "type": "string"
                },
                "ResponseParameters": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "ResponseTemplates": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "SelectionPattern": {
                    "type": "string"
                },
                "StatusCode": {
                    "type": "string"
                }
            },
            "required": [
                "StatusCode"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Method.MethodResponse": {
            "additionalProperties": false,
            "properties": {
                "ResponseModels": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "ResponseParameters": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "boolean"
                        }
                    },
                    "type": "object"
                },
                "StatusCode": {
                    "type": "string"
                }
            },
            "required": [
                "StatusCode"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Model": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ContentType": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "Schema": {
                            "type": "object"
                        }
                    },
                    "required": [
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Model"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::RequestValidator": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "ValidateRequestBody": {
                            "type": "boolean"
                        },
                        "ValidateRequestParameters": {
                            "type": "boolean"
                        }
                    },
                    "required": [
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::RequestValidator"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Resource": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ParentId": {
                            "type": "string"
                        },
                        "PathPart": {
                            "type": "string"
                        },
                        "RestApiId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ParentId",
                        "PathPart",
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Resource"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::RestApi": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApiKeySourceType": {
                            "type": "string"
                        },
                        "BinaryMediaTypes": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Body": {
                            "type": "object"
                        },
                        "BodyS3Location": {
                            "$ref": "#/definitions/AWS::ApiGateway::RestApi.S3Location"
                        },
                        "CloneFrom": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "EndpointConfiguration": {
                            "$ref": "#/definitions/AWS::ApiGateway::RestApi.EndpointConfiguration"
                        },
                        "FailOnWarnings": {
                            "type": "boolean"
                        },
                        "MinimumCompressionSize": {
                            "type": "number"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Parameters": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::RestApi"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::RestApi.EndpointConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Types": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::RestApi.S3Location": {
            "additionalProperties": false,
            "properties": {
                "Bucket": {
                    "type": "string"
                },
                "ETag": {
                    "type": "string"
                },
                "Key": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::Stage": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CacheClusterEnabled": {
                            "type": "boolean"
                        },
                        "CacheClusterSize": {
                            "type": "string"
                        },
                        "ClientCertificateId": {
                            "type": "string"
                        },
                        "DeploymentId": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "DocumentationVersion": {
                            "type": "string"
                        },
                        "MethodSettings": {
                            "items": {
                                "$ref": "#/definitions/AWS::ApiGateway::Stage.MethodSetting"
                            },
                            "type": "array"
                        },
                        "RestApiId": {
                            "type": "string"
                        },
                        "StageName": {
                            "type": "string"
                        },
                        "Variables": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "required": [
                        "RestApiId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::Stage"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::Stage.MethodSetting": {
            "additionalProperties": false,
            "properties": {
                "CacheDataEncrypted": {
                    "type": "boolean"
                },
                "CacheTtlInSeconds": {
                    "type": "number"
                },
                "CachingEnabled": {
                    "type": "boolean"
                },
                "DataTraceEnabled": {
                    "type": "boolean"
                },
                "HttpMethod": {
                    "type": "string"
                },
                "LoggingLevel": {
                    "type": "string"
                },
                "MetricsEnabled": {
                    "type": "boolean"
                },
                "ResourcePath": {
                    "type": "string"
                },
                "ThrottlingBurstLimit": {
                    "type": "number"
                },
                "ThrottlingRateLimit": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::UsagePlan": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApiStages": {
                            "items": {
                                "$ref": "#/definitions/AWS::ApiGateway::UsagePlan.ApiStage"
                            },
                            "type": "array"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Quota": {
                            "$ref": "#/definitions/AWS::ApiGateway::UsagePlan.QuotaSettings"
                        },
                        "Throttle": {
                            "$ref": "#/definitions/AWS::ApiGateway::UsagePlan.ThrottleSettings"
                        },
                        "UsagePlanName": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::UsagePlan"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::UsagePlan.ApiStage": {
            "additionalProperties": false,
            "properties": {
                "ApiId": {
                    "type": "string"
                },
                "Stage": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::UsagePlan.QuotaSettings": {
            "additionalProperties": false,
            "properties": {
                "Limit": {
                    "type": "number"
                },
                "Offset": {
                    "type": "number"
                },
                "Period": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::UsagePlan.ThrottleSettings": {
            "additionalProperties": false,
            "properties": {
                "BurstLimit": {
                    "type": "number"
                },
                "RateLimit": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ApiGateway::UsagePlanKey": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "KeyId": {
                            "type": "string"
                        },
                        "KeyType": {
                            "type": "string"
                        },
                        "UsagePlanId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "KeyId",
                        "KeyType",
                        "UsagePlanId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::UsagePlanKey"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApiGateway::VpcLink": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "TargetArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name",
                        "TargetArns"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApiGateway::VpcLink"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalableTarget": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "MaxCapacity": {
                            "type": "number"
                        },
                        "MinCapacity": {
                            "type": "number"
                        },
                        "ResourceId": {
                            "type": "string"
                        },
                        "RoleARN": {
                            "type": "string"
                        },
                        "ScalableDimension": {
                            "type": "string"
                        },
                        "ScheduledActions": {
                            "items": {
                                "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalableTarget.ScheduledAction"
                            },
                            "type": "array"
                        },
                        "ServiceNamespace": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "MaxCapacity",
                        "MinCapacity",
                        "ResourceId",
                        "RoleARN",
                        "ScalableDimension",
                        "ServiceNamespace"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApplicationAutoScaling::ScalableTarget"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalableTarget.ScalableTargetAction": {
            "additionalProperties": false,
            "properties": {
                "MaxCapacity": {
                    "type": "number"
                },
                "MinCapacity": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalableTarget.ScheduledAction": {
            "additionalProperties": false,
            "properties": {
                "EndTime": {
                    "type": "string"
                },
                "ScalableTargetAction": {
                    "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalableTarget.ScalableTargetAction"
                },
                "Schedule": {
                    "type": "string"
                },
                "ScheduledActionName": {
                    "type": "string"
                },
                "StartTime": {
                    "type": "string"
                }
            },
            "required": [
                "Schedule",
                "ScheduledActionName"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PolicyName": {
                            "type": "string"
                        },
                        "PolicyType": {
                            "type": "string"
                        },
                        "ResourceId": {
                            "type": "string"
                        },
                        "ScalableDimension": {
                            "type": "string"
                        },
                        "ScalingTargetId": {
                            "type": "string"
                        },
                        "ServiceNamespace": {
                            "type": "string"
                        },
                        "StepScalingPolicyConfiguration": {
                            "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy.StepScalingPolicyConfiguration"
                        },
                        "TargetTrackingScalingPolicyConfiguration": {
                            "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingScalingPolicyConfiguration"
                        }
                    },
                    "required": [
                        "PolicyName",
                        "PolicyType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ApplicationAutoScaling::ScalingPolicy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy.CustomizedMetricSpecification": {
            "additionalProperties": false,
            "properties": {
                "Dimensions": {
                    "items": {
                        "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension"
                    },
                    "type": "array"
                },
                "MetricName": {
                    "type": "string"
                },
                "Namespace": {
                    "type": "string"
                },
                "Statistic": {
                    "type": "string"
                },
                "Unit": {
                    "type": "string"
                }
            },
            "required": [
                "MetricName",
                "Namespace",
                "Statistic"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Value"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification": {
            "additionalProperties": false,
            "properties": {
                "PredefinedMetricType": {
                    "type": "string"
                },
                "ResourceLabel": {
                    "type": "string"
                }
            },
            "required": [
                "PredefinedMetricType"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment": {
            "additionalProperties": false,
            "properties": {
                "MetricIntervalLowerBound": {
                    "type": "number"
                },
                "MetricIntervalUpperBound": {
                    "type": "number"
                },
                "ScalingAdjustment": {
                    "type": "number"
                }
            },
            "required": [
                "ScalingAdjustment"
            ],
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy.StepScalingPolicyConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AdjustmentType": {
                    "type": "string"
                },
                "Cooldown": {
                    "type": "number"
                },
                "MetricAggregationType": {
                    "type": "string"
                },
                "MinAdjustmentMagnitude": {
                    "type": "number"
                },
                "StepAdjustments": {
                    "items": {
                        "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingScalingPolicyConfiguration": {
            "additionalProperties": false,
            "properties": {
                "CustomizedMetricSpecification": {
                    "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy.CustomizedMetricSpecification"
                },
                "DisableScaleIn": {
                    "type": "boolean"
                },
                "PredefinedMetricSpecification": {
                    "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification"
                },
                "ScaleInCooldown": {
                    "type": "number"
                },
                "ScaleOutCooldown": {
                    "type": "number"
                },
                "TargetValue": {
                    "type": "number"
                }
            },
            "required": [
                "TargetValue"
            ],
            "type": "object"
        },
        "AWS::Athena::NamedQuery": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Database": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "QueryString": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Database",
                        "QueryString"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Athena::NamedQuery"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::AutoScalingGroup": {
            "additionalProperties": false,
            "properties": {
                "CreationPolicy": {
                    "type": "object"
                },
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutoScalingGroupName": {
                            "type": "string"
                        },
                        "AvailabilityZones": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Cooldown": {
                            "type": "string"
                        },
                        "DesiredCapacity": {
                            "type": "string"
                        },
                        "HealthCheckGracePeriod": {
                            "type": "number"
                        },
                        "HealthCheckType": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "LaunchConfigurationName": {
                            "type": "string"
                        },
                        "LifecycleHookSpecificationList": {
                            "items": {
                                "$ref": "#/definitions/AWS::AutoScaling::AutoScalingGroup.LifecycleHookSpecification"
                            },
                            "type": "array"
                        },
                        "LoadBalancerNames": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "MaxSize": {
                            "type": "string"
                        },
                        "MetricsCollection": {
                            "items": {
                                "$ref": "#/definitions/AWS::AutoScaling::AutoScalingGroup.MetricsCollection"
                            },
                            "type": "array"
                        },
                        "MinSize": {
                            "type": "string"
                        },
                        "NotificationConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::AutoScaling::AutoScalingGroup.NotificationConfiguration"
                            },
                            "type": "array"
                        },
                        "PlacementGroup": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/AWS::AutoScaling::AutoScalingGroup.TagProperty"
                            },
                            "type": "array"
                        },
                        "TargetGroupARNs": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "TerminationPolicies": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "VPCZoneIdentifier": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "MaxSize",
                        "MinSize"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::AutoScaling::AutoScalingGroup"
                    ],
                    "type": "string"
                },
                "UpdatePolicy": {
                    "type": "object"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::AutoScalingGroup.LifecycleHookSpecification": {
            "additionalProperties": false,
            "properties": {
                "DefaultResult": {
                    "type": "string"
                },
                "HeartbeatTimeout": {
                    "type": "number"
                },
                "LifecycleHookName": {
                    "type": "string"
                },
                "LifecycleTransition": {
                    "type": "string"
                },
                "NotificationMetadata": {
                    "type": "string"
                },
                "NotificationTargetARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "LifecycleHookName",
                "LifecycleTransition"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::AutoScalingGroup.MetricsCollection": {
            "additionalProperties": false,
            "properties": {
                "Granularity": {
                    "type": "string"
                },
                "Metrics": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Granularity"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::AutoScalingGroup.NotificationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "NotificationTypes": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "TopicARN": {
                    "type": "string"
                }
            },
            "required": [
                "TopicARN"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::AutoScalingGroup.TagProperty": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "PropagateAtLaunch": {
                    "type": "boolean"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "PropagateAtLaunch",
                "Value"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::LaunchConfiguration": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AssociatePublicIpAddress": {
                            "type": "boolean"
                        },
                        "BlockDeviceMappings": {
                            "items": {
                                "$ref": "#/definitions/AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping"
                            },
                            "type": "array"
                        },
                        "ClassicLinkVPCId": {
                            "type": "string"
                        },
                        "ClassicLinkVPCSecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "EbsOptimized": {
                            "type": "boolean"
                        },
                        "IamInstanceProfile": {
                            "type": "string"
                        },
                        "ImageId": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "InstanceMonitoring": {
                            "type": "boolean"
                        },
                        "InstanceType": {
                            "type": "string"
                        },
                        "KernelId": {
                            "type": "string"
                        },
                        "KeyName": {
                            "type": "string"
                        },
                        "PlacementTenancy": {
                            "type": "string"
                        },
                        "RamDiskId": {
                            "type": "string"
                        },
                        "SecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SpotPrice": {
                            "type": "string"
                        },
                        "UserData": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ImageId",
                        "InstanceType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::AutoScaling::LaunchConfiguration"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::LaunchConfiguration.BlockDevice": {
            "additionalProperties": false,
            "properties": {
                "DeleteOnTermination": {
                    "type": "boolean"
                },
                "Encrypted": {
                    "type": "boolean"
                },
                "Iops": {
                    "type": "number"
                },
                "SnapshotId": {
                    "type": "string"
                },
                "VolumeSize": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping": {
            "additionalProperties": false,
            "properties": {
                "DeviceName": {
                    "type": "string"
                },
                "Ebs": {
                    "$ref": "#/definitions/AWS::AutoScaling::LaunchConfiguration.BlockDevice"
                },
                "NoDevice": {
                    "type": "boolean"
                },
                "VirtualName": {
                    "type": "string"
                }
            },
            "required": [
                "DeviceName"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::LifecycleHook": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutoScalingGroupName": {
                            "type": "string"
                        },
                        "DefaultResult": {
                            "type": "string"
                        },
                        "HeartbeatTimeout": {
                            "type": "number"
                        },
                        "LifecycleHookName": {
                            "type": "string"
                        },
                        "LifecycleTransition": {
                            "type": "string"
                        },
                        "NotificationMetadata": {
                            "type": "string"
                        },
                        "NotificationTargetARN": {
                            "type": "string"
                        },
                        "RoleARN": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AutoScalingGroupName",
                        "LifecycleTransition"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::AutoScaling::LifecycleHook"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScalingPolicy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AdjustmentType": {
                            "type": "string"
                        },
                        "AutoScalingGroupName": {
                            "type": "string"
                        },
                        "Cooldown": {
                            "type": "string"
                        },
                        "EstimatedInstanceWarmup": {
                            "type": "number"
                        },
                        "MetricAggregationType": {
                            "type": "string"
                        },
                        "MinAdjustmentMagnitude": {
                            "type": "number"
                        },
                        "PolicyType": {
                            "type": "string"
                        },
                        "ScalingAdjustment": {
                            "type": "number"
                        },
                        "StepAdjustments": {
                            "items": {
                                "$ref": "#/definitions/AWS::AutoScaling::ScalingPolicy.StepAdjustment"
                            },
                            "type": "array"
                        },
                        "TargetTrackingConfiguration": {
                            "$ref": "#/definitions/AWS::AutoScaling::ScalingPolicy.TargetTrackingConfiguration"
                        }
                    },
                    "required": [
                        "AutoScalingGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::AutoScaling::ScalingPolicy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScalingPolicy.CustomizedMetricSpecification": {
            "additionalProperties": false,
            "properties": {
                "Dimensions": {
                    "items": {
                        "$ref": "#/definitions/AWS::AutoScaling::ScalingPolicy.MetricDimension"
                    },
                    "type": "array"
                },
                "MetricName": {
                    "type": "string"
                },
                "Namespace": {
                    "type": "string"
                },
                "Statistic": {
                    "type": "string"
                },
                "Unit": {
                    "type": "string"
                }
            },
            "required": [
                "MetricName",
                "Namespace",
                "Statistic"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScalingPolicy.MetricDimension": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Value"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScalingPolicy.PredefinedMetricSpecification": {
            "additionalProperties": false,
            "properties": {
                "PredefinedMetricType": {
                    "type": "string"
                },
                "ResourceLabel": {
                    "type": "string"
                }
            },
            "required": [
                "PredefinedMetricType"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScalingPolicy.StepAdjustment": {
            "additionalProperties": false,
            "properties": {
                "MetricIntervalLowerBound": {
                    "type": "number"
                },
                "MetricIntervalUpperBound": {
                    "type": "number"
                },
                "ScalingAdjustment": {
                    "type": "number"
                }
            },
            "required": [
                "ScalingAdjustment"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScalingPolicy.TargetTrackingConfiguration": {
            "additionalProperties": false,
            "properties": {
                "CustomizedMetricSpecification": {
                    "$ref": "#/definitions/AWS::AutoScaling::ScalingPolicy.CustomizedMetricSpecification"
                },
                "DisableScaleIn": {
                    "type": "boolean"
                },
                "PredefinedMetricSpecification": {
                    "$ref": "#/definitions/AWS::AutoScaling::ScalingPolicy.PredefinedMetricSpecification"
                },
                "TargetValue": {
                    "type": "number"
                }
            },
            "required": [
                "TargetValue"
            ],
            "type": "object"
        },
        "AWS::AutoScaling::ScheduledAction": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutoScalingGroupName": {
                            "type": "string"
                        },
                        "DesiredCapacity": {
                            "type": "number"
                        },
                        "EndTime": {
                            "type": "string"
                        },
                        "MaxSize": {
                            "type": "number"
                        },
                        "MinSize": {
                            "type": "number"
                        },
                        "Recurrence": {
                            "type": "string"
                        },
                        "StartTime": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AutoScalingGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::AutoScaling::ScheduledAction"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Batch::ComputeEnvironment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ComputeEnvironmentName": {
                            "type": "string"
                        },
                        "ComputeResources": {
                            "$ref": "#/definitions/AWS::Batch::ComputeEnvironment.ComputeResources"
                        },
                        "ServiceRole": {
                            "type": "string"
                        },
                        "State": {
                            "type": "string"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ComputeResources",
                        "ServiceRole",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Batch::ComputeEnvironment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Batch::ComputeEnvironment.ComputeResources": {
            "additionalProperties": false,
            "properties": {
                "BidPercentage": {
                    "type": "number"
                },
                "DesiredvCpus": {
                    "type": "number"
                },
                "Ec2KeyPair": {
                    "type": "string"
                },
                "ImageId": {
                    "type": "string"
                },
                "InstanceRole": {
                    "type": "string"
                },
                "InstanceTypes": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "MaxvCpus": {
                    "type": "number"
                },
                "MinvCpus": {
                    "type": "number"
                },
                "SecurityGroupIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SpotIamFleetRole": {
                    "type": "string"
                },
                "Subnets": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Tags": {
                    "type": "object"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "InstanceRole",
                "InstanceTypes",
                "MaxvCpus",
                "MinvCpus",
                "SecurityGroupIds",
                "Subnets",
                "Type"
            ],
            "type": "object"
        },
        "AWS::Batch::JobDefinition": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ContainerProperties": {
                            "$ref": "#/definitions/AWS::Batch::JobDefinition.ContainerProperties"
                        },
                        "JobDefinitionName": {
                            "type": "string"
                        },
                        "Parameters": {
                            "type": "object"
                        },
                        "RetryStrategy": {
                            "$ref": "#/definitions/AWS::Batch::JobDefinition.RetryStrategy"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ContainerProperties",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Batch::JobDefinition"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Batch::JobDefinition.ContainerProperties": {
            "additionalProperties": false,
            "properties": {
                "Command": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Environment": {
                    "items": {
                        "$ref": "#/definitions/AWS::Batch::JobDefinition.Environment"
                    },
                    "type": "array"
                },
                "Image": {
                    "type": "string"
                },
                "JobRoleArn": {
                    "type": "string"
                },
                "Memory": {
                    "type": "number"
                },
                "MountPoints": {
                    "items": {
                        "$ref": "#/definitions/AWS::Batch::JobDefinition.MountPoints"
                    },
                    "type": "array"
                },
                "Privileged": {
                    "type": "boolean"
                },
                "ReadonlyRootFilesystem": {
                    "type": "boolean"
                },
                "Ulimits": {
                    "items": {
                        "$ref": "#/definitions/AWS::Batch::JobDefinition.Ulimit"
                    },
                    "type": "array"
                },
                "User": {
                    "type": "string"
                },
                "Vcpus": {
                    "type": "number"
                },
                "Volumes": {
                    "items": {
                        "$ref": "#/definitions/AWS::Batch::JobDefinition.Volumes"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Image",
                "Memory",
                "Vcpus"
            ],
            "type": "object"
        },
        "AWS::Batch::JobDefinition.Environment": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Batch::JobDefinition.MountPoints": {
            "additionalProperties": false,
            "properties": {
                "ContainerPath": {
                    "type": "string"
                },
                "ReadOnly": {
                    "type": "boolean"
                },
                "SourceVolume": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Batch::JobDefinition.RetryStrategy": {
            "additionalProperties": false,
            "properties": {
                "Attempts": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::Batch::JobDefinition.Ulimit": {
            "additionalProperties": false,
            "properties": {
                "HardLimit": {
                    "type": "number"
                },
                "Name": {
                    "type": "string"
                },
                "SoftLimit": {
                    "type": "number"
                }
            },
            "required": [
                "HardLimit",
                "Name",
                "SoftLimit"
            ],
            "type": "object"
        },
        "AWS::Batch::JobDefinition.Volumes": {
            "additionalProperties": false,
            "properties": {
                "Host": {
                    "$ref": "#/definitions/AWS::Batch::JobDefinition.VolumesHost"
                },
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Batch::JobDefinition.VolumesHost": {
            "additionalProperties": false,
            "properties": {
                "SourcePath": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Batch::JobQueue": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ComputeEnvironmentOrder": {
                            "items": {
                                "$ref": "#/definitions/AWS::Batch::JobQueue.ComputeEnvironmentOrder"
                            },
                            "type": "array"
                        },
                        "JobQueueName": {
                            "type": "string"
                        },
                        "Priority": {
                            "type": "number"
                        },
                        "State": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ComputeEnvironmentOrder",
                        "Priority"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Batch::JobQueue"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Batch::JobQueue.ComputeEnvironmentOrder": {
            "additionalProperties": false,
            "properties": {
                "ComputeEnvironment": {
                    "type": "string"
                },
                "Order": {
                    "type": "number"
                }
            },
            "required": [
                "ComputeEnvironment",
                "Order"
            ],
            "type": "object"
        },
        "AWS::CertificateManager::Certificate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DomainName": {
                            "type": "string"
                        },
                        "DomainValidationOptions": {
                            "items": {
                                "$ref": "#/definitions/AWS::CertificateManager::Certificate.DomainValidationOption"
                            },
                            "type": "array"
                        },
                        "SubjectAlternativeNames": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DomainName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CertificateManager::Certificate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CertificateManager::Certificate.DomainValidationOption": {
            "additionalProperties": false,
            "properties": {
                "DomainName": {
                    "type": "string"
                },
                "ValidationDomain": {
                    "type": "string"
                }
            },
            "required": [
                "DomainName",
                "ValidationDomain"
            ],
            "type": "object"
        },
        "AWS::Cloud9::EnvironmentEC2": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutomaticStopTimeMinutes": {
                            "type": "number"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "InstanceType": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "OwnerArn": {
                            "type": "string"
                        },
                        "Repositories": {
                            "items": {
                                "$ref": "#/definitions/AWS::Cloud9::EnvironmentEC2.Repository"
                            },
                            "type": "array"
                        },
                        "SubnetId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "InstanceType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cloud9::EnvironmentEC2"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Cloud9::EnvironmentEC2.Repository": {
            "additionalProperties": false,
            "properties": {
                "PathComponent": {
                    "type": "string"
                },
                "RepositoryUrl": {
                    "type": "string"
                }
            },
            "required": [
                "PathComponent",
                "RepositoryUrl"
            ],
            "type": "object"
        },
        "AWS::CloudFormation::CustomResource": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ServiceToken": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ServiceToken"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFormation::CustomResource"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudFormation::Stack": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "NotificationARNs": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Parameters": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TemplateURL": {
                            "type": "string"
                        },
                        "TimeoutInMinutes": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "TemplateURL"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFormation::Stack"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudFormation::WaitCondition": {
            "additionalProperties": false,
            "properties": {
                "CreationPolicy": {
                    "type": "object"
                },
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Count": {
                            "type": "number"
                        },
                        "Handle": {
                            "type": "string"
                        },
                        "Timeout": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Handle",
                        "Timeout"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFormation::WaitCondition"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudFormation::WaitConditionHandle": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {},
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFormation::WaitConditionHandle"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CloudFront::CloudFrontOriginAccessIdentity": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CloudFrontOriginAccessIdentityConfig": {
                            "$ref": "#/definitions/AWS::CloudFront::CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig"
                        }
                    },
                    "required": [
                        "CloudFrontOriginAccessIdentityConfig"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFront::CloudFrontOriginAccessIdentity"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudFront::CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig": {
            "additionalProperties": false,
            "properties": {
                "Comment": {
                    "type": "string"
                }
            },
            "required": [
                "Comment"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DistributionConfig": {
                            "$ref": "#/definitions/AWS::CloudFront::Distribution.DistributionConfig"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DistributionConfig"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFront::Distribution"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.CacheBehavior": {
            "additionalProperties": false,
            "properties": {
                "AllowedMethods": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "CachedMethods": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Compress": {
                    "type": "boolean"
                },
                "DefaultTTL": {
                    "type": "number"
                },
                "ForwardedValues": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.ForwardedValues"
                },
                "LambdaFunctionAssociations": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudFront::Distribution.LambdaFunctionAssociation"
                    },
                    "type": "array"
                },
                "MaxTTL": {
                    "type": "number"
                },
                "MinTTL": {
                    "type": "number"
                },
                "PathPattern": {
                    "type": "string"
                },
                "SmoothStreaming": {
                    "type": "boolean"
                },
                "TargetOriginId": {
                    "type": "string"
                },
                "TrustedSigners": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "ViewerProtocolPolicy": {
                    "type": "string"
                }
            },
            "required": [
                "ForwardedValues",
                "PathPattern",
                "TargetOriginId",
                "ViewerProtocolPolicy"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.Cookies": {
            "additionalProperties": false,
            "properties": {
                "Forward": {
                    "type": "string"
                },
                "WhitelistedNames": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Forward"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.CustomErrorResponse": {
            "additionalProperties": false,
            "properties": {
                "ErrorCachingMinTTL": {
                    "type": "number"
                },
                "ErrorCode": {
                    "type": "number"
                },
                "ResponseCode": {
                    "type": "number"
                },
                "ResponsePagePath": {
                    "type": "string"
                }
            },
            "required": [
                "ErrorCode"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.CustomOriginConfig": {
            "additionalProperties": false,
            "properties": {
                "HTTPPort": {
                    "type": "number"
                },
                "HTTPSPort": {
                    "type": "number"
                },
                "OriginKeepaliveTimeout": {
                    "type": "number"
                },
                "OriginProtocolPolicy": {
                    "type": "string"
                },
                "OriginReadTimeout": {
                    "type": "number"
                },
                "OriginSSLProtocols": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "OriginProtocolPolicy"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.DefaultCacheBehavior": {
            "additionalProperties": false,
            "properties": {
                "AllowedMethods": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "CachedMethods": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Compress": {
                    "type": "boolean"
                },
                "DefaultTTL": {
                    "type": "number"
                },
                "ForwardedValues": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.ForwardedValues"
                },
                "LambdaFunctionAssociations": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudFront::Distribution.LambdaFunctionAssociation"
                    },
                    "type": "array"
                },
                "MaxTTL": {
                    "type": "number"
                },
                "MinTTL": {
                    "type": "number"
                },
                "SmoothStreaming": {
                    "type": "boolean"
                },
                "TargetOriginId": {
                    "type": "string"
                },
                "TrustedSigners": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "ViewerProtocolPolicy": {
                    "type": "string"
                }
            },
            "required": [
                "ForwardedValues",
                "TargetOriginId",
                "ViewerProtocolPolicy"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.DistributionConfig": {
            "additionalProperties": false,
            "properties": {
                "Aliases": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "CacheBehaviors": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudFront::Distribution.CacheBehavior"
                    },
                    "type": "array"
                },
                "Comment": {
                    "type": "string"
                },
                "CustomErrorResponses": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudFront::Distribution.CustomErrorResponse"
                    },
                    "type": "array"
                },
                "DefaultCacheBehavior": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.DefaultCacheBehavior"
                },
                "DefaultRootObject": {
                    "type": "string"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "HttpVersion": {
                    "type": "string"
                },
                "IPV6Enabled": {
                    "type": "boolean"
                },
                "Logging": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.Logging"
                },
                "Origins": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudFront::Distribution.Origin"
                    },
                    "type": "array"
                },
                "PriceClass": {
                    "type": "string"
                },
                "Restrictions": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.Restrictions"
                },
                "ViewerCertificate": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.ViewerCertificate"
                },
                "WebACLId": {
                    "type": "string"
                }
            },
            "required": [
                "Enabled"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.ForwardedValues": {
            "additionalProperties": false,
            "properties": {
                "Cookies": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.Cookies"
                },
                "Headers": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "QueryString": {
                    "type": "boolean"
                },
                "QueryStringCacheKeys": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "QueryString"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.GeoRestriction": {
            "additionalProperties": false,
            "properties": {
                "Locations": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "RestrictionType": {
                    "type": "string"
                }
            },
            "required": [
                "RestrictionType"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.LambdaFunctionAssociation": {
            "additionalProperties": false,
            "properties": {
                "EventType": {
                    "type": "string"
                },
                "LambdaFunctionARN": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CloudFront::Distribution.Logging": {
            "additionalProperties": false,
            "properties": {
                "Bucket": {
                    "type": "string"
                },
                "IncludeCookies": {
                    "type": "boolean"
                },
                "Prefix": {
                    "type": "string"
                }
            },
            "required": [
                "Bucket"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.Origin": {
            "additionalProperties": false,
            "properties": {
                "CustomOriginConfig": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.CustomOriginConfig"
                },
                "DomainName": {
                    "type": "string"
                },
                "Id": {
                    "type": "string"
                },
                "OriginCustomHeaders": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudFront::Distribution.OriginCustomHeader"
                    },
                    "type": "array"
                },
                "OriginPath": {
                    "type": "string"
                },
                "S3OriginConfig": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.S3OriginConfig"
                }
            },
            "required": [
                "DomainName",
                "Id"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.OriginCustomHeader": {
            "additionalProperties": false,
            "properties": {
                "HeaderName": {
                    "type": "string"
                },
                "HeaderValue": {
                    "type": "string"
                }
            },
            "required": [
                "HeaderName",
                "HeaderValue"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.Restrictions": {
            "additionalProperties": false,
            "properties": {
                "GeoRestriction": {
                    "$ref": "#/definitions/AWS::CloudFront::Distribution.GeoRestriction"
                }
            },
            "required": [
                "GeoRestriction"
            ],
            "type": "object"
        },
        "AWS::CloudFront::Distribution.S3OriginConfig": {
            "additionalProperties": false,
            "properties": {
                "OriginAccessIdentity": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CloudFront::Distribution.ViewerCertificate": {
            "additionalProperties": false,
            "properties": {
                "AcmCertificateArn": {
                    "type": "string"
                },
                "CloudFrontDefaultCertificate": {
                    "type": "boolean"
                },
                "IamCertificateId": {
                    "type": "string"
                },
                "MinimumProtocolVersion": {
                    "type": "string"
                },
                "SslSupportMethod": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CloudFront::StreamingDistribution": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "StreamingDistributionConfig": {
                            "$ref": "#/definitions/AWS::CloudFront::StreamingDistribution.StreamingDistributionConfig"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "StreamingDistributionConfig",
                        "Tags"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudFront::StreamingDistribution"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudFront::StreamingDistribution.Logging": {
            "additionalProperties": false,
            "properties": {
                "Bucket": {
                    "type": "string"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "Prefix": {
                    "type": "string"
                }
            },
            "required": [
                "Bucket",
                "Enabled",
                "Prefix"
            ],
            "type": "object"
        },
        "AWS::CloudFront::StreamingDistribution.S3Origin": {
            "additionalProperties": false,
            "properties": {
                "DomainName": {
                    "type": "string"
                },
                "OriginAccessIdentity": {
                    "type": "string"
                }
            },
            "required": [
                "DomainName",
                "OriginAccessIdentity"
            ],
            "type": "object"
        },
        "AWS::CloudFront::StreamingDistribution.StreamingDistributionConfig": {
            "additionalProperties": false,
            "properties": {
                "Aliases": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Comment": {
                    "type": "string"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "Logging": {
                    "$ref": "#/definitions/AWS::CloudFront::StreamingDistribution.Logging"
                },
                "PriceClass": {
                    "type": "string"
                },
                "S3Origin": {
                    "$ref": "#/definitions/AWS::CloudFront::StreamingDistribution.S3Origin"
                },
                "TrustedSigners": {
                    "$ref": "#/definitions/AWS::CloudFront::StreamingDistribution.TrustedSigners"
                }
            },
            "required": [
                "Comment",
                "Enabled",
                "S3Origin",
                "TrustedSigners"
            ],
            "type": "object"
        },
        "AWS::CloudFront::StreamingDistribution.TrustedSigners": {
            "additionalProperties": false,
            "properties": {
                "AwsAccountNumbers": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Enabled": {
                    "type": "boolean"
                }
            },
            "required": [
                "Enabled"
            ],
            "type": "object"
        },
        "AWS::CloudTrail::Trail": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CloudWatchLogsLogGroupArn": {
                            "type": "string"
                        },
                        "CloudWatchLogsRoleArn": {
                            "type": "string"
                        },
                        "EnableLogFileValidation": {
                            "type": "boolean"
                        },
                        "EventSelectors": {
                            "items": {
                                "$ref": "#/definitions/AWS::CloudTrail::Trail.EventSelector"
                            },
                            "type": "array"
                        },
                        "IncludeGlobalServiceEvents": {
                            "type": "boolean"
                        },
                        "IsLogging": {
                            "type": "boolean"
                        },
                        "IsMultiRegionTrail": {
                            "type": "boolean"
                        },
                        "KMSKeyId": {
                            "type": "string"
                        },
                        "S3BucketName": {
                            "type": "string"
                        },
                        "S3KeyPrefix": {
                            "type": "string"
                        },
                        "SnsTopicName": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TrailName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "IsLogging",
                        "S3BucketName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudTrail::Trail"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudTrail::Trail.DataResource": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CloudTrail::Trail.EventSelector": {
            "additionalProperties": false,
            "properties": {
                "DataResources": {
                    "items": {
                        "$ref": "#/definitions/AWS::CloudTrail::Trail.DataResource"
                    },
                    "type": "array"
                },
                "IncludeManagementEvents": {
                    "type": "boolean"
                },
                "ReadWriteType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CloudWatch::Alarm": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ActionsEnabled": {
                            "type": "boolean"
                        },
                        "AlarmActions": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "AlarmDescription": {
                            "type": "string"
                        },
                        "AlarmName": {
                            "type": "string"
                        },
                        "ComparisonOperator": {
                            "type": "string"
                        },
                        "Dimensions": {
                            "items": {
                                "$ref": "#/definitions/AWS::CloudWatch::Alarm.Dimension"
                            },
                            "type": "array"
                        },
                        "EvaluateLowSampleCountPercentile": {
                            "type": "string"
                        },
                        "EvaluationPeriods": {
                            "type": "number"
                        },
                        "ExtendedStatistic": {
                            "type": "string"
                        },
                        "InsufficientDataActions": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "MetricName": {
                            "type": "string"
                        },
                        "Namespace": {
                            "type": "string"
                        },
                        "OKActions": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Period": {
                            "type": "number"
                        },
                        "Statistic": {
                            "type": "string"
                        },
                        "Threshold": {
                            "type": "number"
                        },
                        "TreatMissingData": {
                            "type": "string"
                        },
                        "Unit": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ComparisonOperator",
                        "EvaluationPeriods",
                        "MetricName",
                        "Namespace",
                        "Period",
                        "Threshold"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudWatch::Alarm"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CloudWatch::Alarm.Dimension": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Value"
            ],
            "type": "object"
        },
        "AWS::CloudWatch::Dashboard": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DashboardBody": {
                            "type": "string"
                        },
                        "DashboardName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DashboardBody"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CloudWatch::Dashboard"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Artifacts": {
                            "$ref": "#/definitions/AWS::CodeBuild::Project.Artifacts"
                        },
                        "BadgeEnabled": {
                            "type": "boolean"
                        },
                        "Cache": {
                            "$ref": "#/definitions/AWS::CodeBuild::Project.ProjectCache"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "EncryptionKey": {
                            "type": "string"
                        },
                        "Environment": {
                            "$ref": "#/definitions/AWS::CodeBuild::Project.Environment"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "ServiceRole": {
                            "type": "string"
                        },
                        "Source": {
                            "$ref": "#/definitions/AWS::CodeBuild::Project.Source"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TimeoutInMinutes": {
                            "type": "number"
                        },
                        "Triggers": {
                            "$ref": "#/definitions/AWS::CodeBuild::Project.ProjectTriggers"
                        },
                        "VpcConfig": {
                            "$ref": "#/definitions/AWS::CodeBuild::Project.VpcConfig"
                        }
                    },
                    "required": [
                        "Artifacts",
                        "Environment",
                        "ServiceRole",
                        "Source"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodeBuild::Project"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.Artifacts": {
            "additionalProperties": false,
            "properties": {
                "Location": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "NamespaceType": {
                    "type": "string"
                },
                "Packaging": {
                    "type": "string"
                },
                "Path": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.Environment": {
            "additionalProperties": false,
            "properties": {
                "ComputeType": {
                    "type": "string"
                },
                "EnvironmentVariables": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodeBuild::Project.EnvironmentVariable"
                    },
                    "type": "array"
                },
                "Image": {
                    "type": "string"
                },
                "PrivilegedMode": {
                    "type": "boolean"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "ComputeType",
                "Image",
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.EnvironmentVariable": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Value"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.ProjectCache": {
            "additionalProperties": false,
            "properties": {
                "Location": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.ProjectTriggers": {
            "additionalProperties": false,
            "properties": {
                "Webhook": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::CodeBuild::Project.Source": {
            "additionalProperties": false,
            "properties": {
                "Auth": {
                    "$ref": "#/definitions/AWS::CodeBuild::Project.SourceAuth"
                },
                "BuildSpec": {
                    "type": "string"
                },
                "GitCloneDepth": {
                    "type": "number"
                },
                "InsecureSsl": {
                    "type": "boolean"
                },
                "Location": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.SourceAuth": {
            "additionalProperties": false,
            "properties": {
                "Resource": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeBuild::Project.VpcConfig": {
            "additionalProperties": false,
            "properties": {
                "SecurityGroupIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Subnets": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "VpcId": {
                    "type": "string"
                }
            },
            "required": [
                "SecurityGroupIds",
                "Subnets",
                "VpcId"
            ],
            "type": "object"
        },
        "AWS::CodeCommit::Repository": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "RepositoryDescription": {
                            "type": "string"
                        },
                        "RepositoryName": {
                            "type": "string"
                        },
                        "Triggers": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodeCommit::Repository.RepositoryTrigger"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "RepositoryName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodeCommit::Repository"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CodeCommit::Repository.RepositoryTrigger": {
            "additionalProperties": false,
            "properties": {
                "Branches": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "CustomData": {
                    "type": "string"
                },
                "DestinationArn": {
                    "type": "string"
                },
                "Events": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::Application": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "ComputePlatform": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodeDeploy::Application"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentConfig": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DeploymentConfigName": {
                            "type": "string"
                        },
                        "MinimumHealthyHosts": {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodeDeploy::DeploymentConfig"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                },
                "Value": {
                    "type": "number"
                }
            },
            "required": [
                "Type",
                "Value"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AlarmConfiguration": {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration"
                        },
                        "ApplicationName": {
                            "type": "string"
                        },
                        "AutoRollbackConfiguration": {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.AutoRollbackConfiguration"
                        },
                        "AutoScalingGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Deployment": {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.Deployment"
                        },
                        "DeploymentConfigName": {
                            "type": "string"
                        },
                        "DeploymentGroupName": {
                            "type": "string"
                        },
                        "DeploymentStyle": {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.DeploymentStyle"
                        },
                        "Ec2TagFilters": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.EC2TagFilter"
                            },
                            "type": "array"
                        },
                        "LoadBalancerInfo": {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.LoadBalancerInfo"
                        },
                        "OnPremisesInstanceTagFilters": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.TagFilter"
                            },
                            "type": "array"
                        },
                        "ServiceRoleArn": {
                            "type": "string"
                        },
                        "TriggerConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.TriggerConfig"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ApplicationName",
                        "ServiceRoleArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodeDeploy::DeploymentGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.Alarm": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Alarms": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.Alarm"
                    },
                    "type": "array"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "IgnorePollAlarmFailure": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.AutoRollbackConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Enabled": {
                    "type": "boolean"
                },
                "Events": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.Deployment": {
            "additionalProperties": false,
            "properties": {
                "Description": {
                    "type": "string"
                },
                "IgnoreApplicationStopFailures": {
                    "type": "boolean"
                },
                "Revision": {
                    "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.RevisionLocation"
                }
            },
            "required": [
                "Revision"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.DeploymentStyle": {
            "additionalProperties": false,
            "properties": {
                "DeploymentOption": {
                    "type": "string"
                },
                "DeploymentType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.EC2TagFilter": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.ELBInfo": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.GitHubLocation": {
            "additionalProperties": false,
            "properties": {
                "CommitId": {
                    "type": "string"
                },
                "Repository": {
                    "type": "string"
                }
            },
            "required": [
                "CommitId",
                "Repository"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.LoadBalancerInfo": {
            "additionalProperties": false,
            "properties": {
                "ElbInfoList": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.ELBInfo"
                    },
                    "type": "array"
                },
                "TargetGroupInfoList": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.TargetGroupInfo"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.RevisionLocation": {
            "additionalProperties": false,
            "properties": {
                "GitHubLocation": {
                    "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.GitHubLocation"
                },
                "RevisionType": {
                    "type": "string"
                },
                "S3Location": {
                    "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup.S3Location"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.S3Location": {
            "additionalProperties": false,
            "properties": {
                "Bucket": {
                    "type": "string"
                },
                "BundleType": {
                    "type": "string"
                },
                "ETag": {
                    "type": "string"
                },
                "Key": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "required": [
                "Bucket",
                "Key"
            ],
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.TagFilter": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.TargetGroupInfo": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodeDeploy::DeploymentGroup.TriggerConfig": {
            "additionalProperties": false,
            "properties": {
                "TriggerEvents": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "TriggerName": {
                    "type": "string"
                },
                "TriggerTargetArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodePipeline::CustomActionType": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Category": {
                            "type": "string"
                        },
                        "ConfigurationProperties": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodePipeline::CustomActionType.ConfigurationProperties"
                            },
                            "type": "array"
                        },
                        "InputArtifactDetails": {
                            "$ref": "#/definitions/AWS::CodePipeline::CustomActionType.ArtifactDetails"
                        },
                        "OutputArtifactDetails": {
                            "$ref": "#/definitions/AWS::CodePipeline::CustomActionType.ArtifactDetails"
                        },
                        "Provider": {
                            "type": "string"
                        },
                        "Settings": {
                            "$ref": "#/definitions/AWS::CodePipeline::CustomActionType.Settings"
                        },
                        "Version": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Category",
                        "InputArtifactDetails",
                        "OutputArtifactDetails",
                        "Provider"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodePipeline::CustomActionType"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::CustomActionType.ArtifactDetails": {
            "additionalProperties": false,
            "properties": {
                "MaximumCount": {
                    "type": "number"
                },
                "MinimumCount": {
                    "type": "number"
                }
            },
            "required": [
                "MaximumCount",
                "MinimumCount"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::CustomActionType.ConfigurationProperties": {
            "additionalProperties": false,
            "properties": {
                "Description": {
                    "type": "string"
                },
                "Key": {
                    "type": "boolean"
                },
                "Name": {
                    "type": "string"
                },
                "Queryable": {
                    "type": "boolean"
                },
                "Required": {
                    "type": "boolean"
                },
                "Secret": {
                    "type": "boolean"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Name",
                "Required",
                "Secret"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::CustomActionType.Settings": {
            "additionalProperties": false,
            "properties": {
                "EntityUrlTemplate": {
                    "type": "string"
                },
                "ExecutionUrlTemplate": {
                    "type": "string"
                },
                "RevisionUrlTemplate": {
                    "type": "string"
                },
                "ThirdPartyConfigurationUrl": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ArtifactStore": {
                            "$ref": "#/definitions/AWS::CodePipeline::Pipeline.ArtifactStore"
                        },
                        "DisableInboundStageTransitions": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodePipeline::Pipeline.StageTransition"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "RestartExecutionOnUpdate": {
                            "type": "boolean"
                        },
                        "RoleArn": {
                            "type": "string"
                        },
                        "Stages": {
                            "items": {
                                "$ref": "#/definitions/AWS::CodePipeline::Pipeline.StageDeclaration"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ArtifactStore",
                        "RoleArn",
                        "Stages"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::CodePipeline::Pipeline"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.ActionDeclaration": {
            "additionalProperties": false,
            "properties": {
                "ActionTypeId": {
                    "$ref": "#/definitions/AWS::CodePipeline::Pipeline.ActionTypeId"
                },
                "Configuration": {
                    "type": "object"
                },
                "InputArtifacts": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodePipeline::Pipeline.InputArtifact"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                },
                "OutputArtifacts": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodePipeline::Pipeline.OutputArtifact"
                    },
                    "type": "array"
                },
                "RoleArn": {
                    "type": "string"
                },
                "RunOrder": {
                    "type": "number"
                }
            },
            "required": [
                "ActionTypeId",
                "Name"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.ActionTypeId": {
            "additionalProperties": false,
            "properties": {
                "Category": {
                    "type": "string"
                },
                "Owner": {
                    "type": "string"
                },
                "Provider": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "required": [
                "Category",
                "Owner",
                "Provider",
                "Version"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.ArtifactStore": {
            "additionalProperties": false,
            "properties": {
                "EncryptionKey": {
                    "$ref": "#/definitions/AWS::CodePipeline::Pipeline.EncryptionKey"
                },
                "Location": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Location",
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.BlockerDeclaration": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.EncryptionKey": {
            "additionalProperties": false,
            "properties": {
                "Id": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Id",
                "Type"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.InputArtifact": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "Name"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.OutputArtifact": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "Name"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.StageDeclaration": {
            "additionalProperties": false,
            "properties": {
                "Actions": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodePipeline::Pipeline.ActionDeclaration"
                    },
                    "type": "array"
                },
                "Blockers": {
                    "items": {
                        "$ref": "#/definitions/AWS::CodePipeline::Pipeline.BlockerDeclaration"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "Actions",
                "Name"
            ],
            "type": "object"
        },
        "AWS::CodePipeline::Pipeline.StageTransition": {
            "additionalProperties": false,
            "properties": {
                "Reason": {
                    "type": "string"
                },
                "StageName": {
                    "type": "string"
                }
            },
            "required": [
                "Reason",
                "StageName"
            ],
            "type": "object"
        },
        "AWS::Cognito::IdentityPool": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllowUnauthenticatedIdentities": {
                            "type": "boolean"
                        },
                        "CognitoEvents": {
                            "type": "object"
                        },
                        "CognitoIdentityProviders": {
                            "items": {
                                "$ref": "#/definitions/AWS::Cognito::IdentityPool.CognitoIdentityProvider"
                            },
                            "type": "array"
                        },
                        "CognitoStreams": {
                            "$ref": "#/definitions/AWS::Cognito::IdentityPool.CognitoStreams"
                        },
                        "DeveloperProviderName": {
                            "type": "string"
                        },
                        "IdentityPoolName": {
                            "type": "string"
                        },
                        "OpenIdConnectProviderARNs": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "PushSync": {
                            "$ref": "#/definitions/AWS::Cognito::IdentityPool.PushSync"
                        },
                        "SamlProviderARNs": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SupportedLoginProviders": {
                            "type": "object"
                        }
                    },
                    "required": [
                        "AllowUnauthenticatedIdentities"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::IdentityPool"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Cognito::IdentityPool.CognitoIdentityProvider": {
            "additionalProperties": false,
            "properties": {
                "ClientId": {
                    "type": "string"
                },
                "ProviderName": {
                    "type": "string"
                },
                "ServerSideTokenCheck": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::IdentityPool.CognitoStreams": {
            "additionalProperties": false,
            "properties": {
                "RoleArn": {
                    "type": "string"
                },
                "StreamName": {
                    "type": "string"
                },
                "StreamingStatus": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::IdentityPool.PushSync": {
            "additionalProperties": false,
            "properties": {
                "ApplicationArns": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "RoleArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::IdentityPoolRoleAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "IdentityPoolId": {
                            "type": "string"
                        },
                        "RoleMappings": {
                            "type": "object"
                        },
                        "Roles": {
                            "type": "object"
                        }
                    },
                    "required": [
                        "IdentityPoolId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::IdentityPoolRoleAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Cognito::IdentityPoolRoleAttachment.MappingRule": {
            "additionalProperties": false,
            "properties": {
                "Claim": {
                    "type": "string"
                },
                "MatchType": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Claim",
                "MatchType",
                "RoleARN",
                "Value"
            ],
            "type": "object"
        },
        "AWS::Cognito::IdentityPoolRoleAttachment.RoleMapping": {
            "additionalProperties": false,
            "properties": {
                "AmbiguousRoleResolution": {
                    "type": "string"
                },
                "RulesConfiguration": {
                    "$ref": "#/definitions/AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType": {
            "additionalProperties": false,
            "properties": {
                "Rules": {
                    "items": {
                        "$ref": "#/definitions/AWS::Cognito::IdentityPoolRoleAttachment.MappingRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Rules"
            ],
            "type": "object"
        },
        "AWS::Cognito::UserPool": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AdminCreateUserConfig": {
                            "$ref": "#/definitions/AWS::Cognito::UserPool.AdminCreateUserConfig"
                        },
                        "AliasAttributes": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "AutoVerifiedAttributes": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "DeviceConfiguration": {
                            "$ref": "#/definitions/AWS::Cognito::UserPool.DeviceConfiguration"
                        },
                        "EmailConfiguration": {
                            "$ref": "#/definitions/AWS::Cognito::UserPool.EmailConfiguration"
                        },
                        "EmailVerificationMessage": {
                            "type": "string"
                        },
                        "EmailVerificationSubject": {
                            "type": "string"
                        },
                        "LambdaConfig": {
                            "$ref": "#/definitions/AWS::Cognito::UserPool.LambdaConfig"
                        },
                        "MfaConfiguration": {
                            "type": "string"
                        },
                        "Policies": {
                            "$ref": "#/definitions/AWS::Cognito::UserPool.Policies"
                        },
                        "Schema": {
                            "items": {
                                "$ref": "#/definitions/AWS::Cognito::UserPool.SchemaAttribute"
                            },
                            "type": "array"
                        },
                        "SmsAuthenticationMessage": {
                            "type": "string"
                        },
                        "SmsConfiguration": {
                            "$ref": "#/definitions/AWS::Cognito::UserPool.SmsConfiguration"
                        },
                        "SmsVerificationMessage": {
                            "type": "string"
                        },
                        "UserPoolName": {
                            "type": "string"
                        },
                        "UserPoolTags": {
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::UserPool"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Cognito::UserPool.AdminCreateUserConfig": {
            "additionalProperties": false,
            "properties": {
                "AllowAdminCreateUserOnly": {
                    "type": "boolean"
                },
                "InviteMessageTemplate": {
                    "$ref": "#/definitions/AWS::Cognito::UserPool.InviteMessageTemplate"
                },
                "UnusedAccountValidityDays": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.DeviceConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ChallengeRequiredOnNewDevice": {
                    "type": "boolean"
                },
                "DeviceOnlyRememberedOnUserPrompt": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.EmailConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ReplyToEmailAddress": {
                    "type": "string"
                },
                "SourceArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.InviteMessageTemplate": {
            "additionalProperties": false,
            "properties": {
                "EmailMessage": {
                    "type": "string"
                },
                "EmailSubject": {
                    "type": "string"
                },
                "SMSMessage": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.LambdaConfig": {
            "additionalProperties": false,
            "properties": {
                "CreateAuthChallenge": {
                    "type": "string"
                },
                "CustomMessage": {
                    "type": "string"
                },
                "DefineAuthChallenge": {
                    "type": "string"
                },
                "PostAuthentication": {
                    "type": "string"
                },
                "PostConfirmation": {
                    "type": "string"
                },
                "PreAuthentication": {
                    "type": "string"
                },
                "PreSignUp": {
                    "type": "string"
                },
                "VerifyAuthChallengeResponse": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.NumberAttributeConstraints": {
            "additionalProperties": false,
            "properties": {
                "MaxValue": {
                    "type": "string"
                },
                "MinValue": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.PasswordPolicy": {
            "additionalProperties": false,
            "properties": {
                "MinimumLength": {
                    "type": "number"
                },
                "RequireLowercase": {
                    "type": "boolean"
                },
                "RequireNumbers": {
                    "type": "boolean"
                },
                "RequireSymbols": {
                    "type": "boolean"
                },
                "RequireUppercase": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.Policies": {
            "additionalProperties": false,
            "properties": {
                "PasswordPolicy": {
                    "$ref": "#/definitions/AWS::Cognito::UserPool.PasswordPolicy"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.SchemaAttribute": {
            "additionalProperties": false,
            "properties": {
                "AttributeDataType": {
                    "type": "string"
                },
                "DeveloperOnlyAttribute": {
                    "type": "boolean"
                },
                "Mutable": {
                    "type": "boolean"
                },
                "Name": {
                    "type": "string"
                },
                "NumberAttributeConstraints": {
                    "$ref": "#/definitions/AWS::Cognito::UserPool.NumberAttributeConstraints"
                },
                "Required": {
                    "type": "boolean"
                },
                "StringAttributeConstraints": {
                    "$ref": "#/definitions/AWS::Cognito::UserPool.StringAttributeConstraints"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.SmsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ExternalId": {
                    "type": "string"
                },
                "SnsCallerArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPool.StringAttributeConstraints": {
            "additionalProperties": false,
            "properties": {
                "MaxLength": {
                    "type": "string"
                },
                "MinLength": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPoolClient": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ClientName": {
                            "type": "string"
                        },
                        "ExplicitAuthFlows": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "GenerateSecret": {
                            "type": "boolean"
                        },
                        "ReadAttributes": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "RefreshTokenValidity": {
                            "type": "number"
                        },
                        "UserPoolId": {
                            "type": "string"
                        },
                        "WriteAttributes": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "UserPoolId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::UserPoolClient"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Cognito::UserPoolGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "GroupName": {
                            "type": "string"
                        },
                        "Precedence": {
                            "type": "number"
                        },
                        "RoleArn": {
                            "type": "string"
                        },
                        "UserPoolId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "UserPoolId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::UserPoolGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Cognito::UserPoolUser": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DesiredDeliveryMediums": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ForceAliasCreation": {
                            "type": "boolean"
                        },
                        "MessageAction": {
                            "type": "string"
                        },
                        "UserAttributes": {
                            "items": {
                                "$ref": "#/definitions/AWS::Cognito::UserPoolUser.AttributeType"
                            },
                            "type": "array"
                        },
                        "UserPoolId": {
                            "type": "string"
                        },
                        "Username": {
                            "type": "string"
                        },
                        "ValidationData": {
                            "items": {
                                "$ref": "#/definitions/AWS::Cognito::UserPoolUser.AttributeType"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "UserPoolId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::UserPoolUser"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Cognito::UserPoolUser.AttributeType": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Cognito::UserPoolUserToGroupAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "GroupName": {
                            "type": "string"
                        },
                        "UserPoolId": {
                            "type": "string"
                        },
                        "Username": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "GroupName",
                        "UserPoolId",
                        "Username"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Cognito::UserPoolUserToGroupAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Config::ConfigRule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ConfigRuleName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "InputParameters": {
                            "type": "object"
                        },
                        "MaximumExecutionFrequency": {
                            "type": "string"
                        },
                        "Scope": {
                            "$ref": "#/definitions/AWS::Config::ConfigRule.Scope"
                        },
                        "Source": {
                            "$ref": "#/definitions/AWS::Config::ConfigRule.Source"
                        }
                    },
                    "required": [
                        "Source"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Config::ConfigRule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Config::ConfigRule.Scope": {
            "additionalProperties": false,
            "properties": {
                "ComplianceResourceId": {
                    "type": "string"
                },
                "ComplianceResourceTypes": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "TagKey": {
                    "type": "string"
                },
                "TagValue": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Config::ConfigRule.Source": {
            "additionalProperties": false,
            "properties": {
                "Owner": {
                    "type": "string"
                },
                "SourceDetails": {
                    "items": {
                        "$ref": "#/definitions/AWS::Config::ConfigRule.SourceDetail"
                    },
                    "type": "array"
                },
                "SourceIdentifier": {
                    "type": "string"
                }
            },
            "required": [
                "Owner",
                "SourceIdentifier"
            ],
            "type": "object"
        },
        "AWS::Config::ConfigRule.SourceDetail": {
            "additionalProperties": false,
            "properties": {
                "EventSource": {
                    "type": "string"
                },
                "MaximumExecutionFrequency": {
                    "type": "string"
                },
                "MessageType": {
                    "type": "string"
                }
            },
            "required": [
                "EventSource",
                "MessageType"
            ],
            "type": "object"
        },
        "AWS::Config::ConfigurationRecorder": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "RecordingGroup": {
                            "$ref": "#/definitions/AWS::Config::ConfigurationRecorder.RecordingGroup"
                        },
                        "RoleARN": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "RoleARN"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Config::ConfigurationRecorder"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Config::ConfigurationRecorder.RecordingGroup": {
            "additionalProperties": false,
            "properties": {
                "AllSupported": {
                    "type": "boolean"
                },
                "IncludeGlobalResourceTypes": {
                    "type": "boolean"
                },
                "ResourceTypes": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::Config::DeliveryChannel": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ConfigSnapshotDeliveryProperties": {
                            "$ref": "#/definitions/AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "S3BucketName": {
                            "type": "string"
                        },
                        "S3KeyPrefix": {
                            "type": "string"
                        },
                        "SnsTopicARN": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "S3BucketName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Config::DeliveryChannel"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties": {
            "additionalProperties": false,
            "properties": {
                "DeliveryFrequency": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::DAX::Cluster": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AvailabilityZones": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ClusterName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "IAMRoleARN": {
                            "type": "string"
                        },
                        "NodeType": {
                            "type": "string"
                        },
                        "NotificationTopicARN": {
                            "type": "string"
                        },
                        "ParameterGroupName": {
                            "type": "string"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "ReplicationFactor": {
                            "type": "number"
                        },
                        "SecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SubnetGroupName": {
                            "type": "string"
                        },
                        "Tags": {
                            "type": "object"
                        }
                    },
                    "required": [
                        "IAMRoleARN",
                        "NodeType",
                        "ReplicationFactor"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DAX::Cluster"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DAX::ParameterGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "ParameterGroupName": {
                            "type": "string"
                        },
                        "ParameterNameValues": {
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DAX::ParameterGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::DAX::SubnetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "SubnetGroupName": {
                            "type": "string"
                        },
                        "SubnetIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "SubnetIds"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DAX::SubnetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DMS::Certificate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CertificateIdentifier": {
                            "type": "string"
                        },
                        "CertificatePem": {
                            "type": "string"
                        },
                        "CertificateWallet": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DMS::Certificate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::DMS::Endpoint": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CertificateArn": {
                            "type": "string"
                        },
                        "DatabaseName": {
                            "type": "string"
                        },
                        "DynamoDbSettings": {
                            "$ref": "#/definitions/AWS::DMS::Endpoint.DynamoDbSettings"
                        },
                        "EndpointIdentifier": {
                            "type": "string"
                        },
                        "EndpointType": {
                            "type": "string"
                        },
                        "EngineName": {
                            "type": "string"
                        },
                        "ExtraConnectionAttributes": {
                            "type": "string"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "MongoDbSettings": {
                            "$ref": "#/definitions/AWS::DMS::Endpoint.MongoDbSettings"
                        },
                        "Password": {
                            "type": "string"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "S3Settings": {
                            "$ref": "#/definitions/AWS::DMS::Endpoint.S3Settings"
                        },
                        "ServerName": {
                            "type": "string"
                        },
                        "SslMode": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Username": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "EndpointType",
                        "EngineName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DMS::Endpoint"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DMS::Endpoint.DynamoDbSettings": {
            "additionalProperties": false,
            "properties": {
                "ServiceAccessRoleArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::DMS::Endpoint.MongoDbSettings": {
            "additionalProperties": false,
            "properties": {
                "AuthMechanism": {
                    "type": "string"
                },
                "AuthSource": {
                    "type": "string"
                },
                "AuthType": {
                    "type": "string"
                },
                "DatabaseName": {
                    "type": "string"
                },
                "DocsToInvestigate": {
                    "type": "string"
                },
                "ExtractDocId": {
                    "type": "string"
                },
                "NestingLevel": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "Port": {
                    "type": "number"
                },
                "ServerName": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::DMS::Endpoint.S3Settings": {
            "additionalProperties": false,
            "properties": {
                "BucketFolder": {
                    "type": "string"
                },
                "BucketName": {
                    "type": "string"
                },
                "CompressionType": {
                    "type": "string"
                },
                "CsvDelimiter": {
                    "type": "string"
                },
                "CsvRowDelimiter": {
                    "type": "string"
                },
                "ExternalTableDefinition": {
                    "type": "string"
                },
                "ServiceAccessRoleArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::DMS::EventSubscription": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Enabled": {
                            "type": "boolean"
                        },
                        "EventCategories": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SnsTopicArn": {
                            "type": "string"
                        },
                        "SourceIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SourceType": {
                            "type": "string"
                        },
                        "SubscriptionName": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "SnsTopicArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DMS::EventSubscription"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DMS::ReplicationInstance": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllocatedStorage": {
                            "type": "number"
                        },
                        "AllowMajorVersionUpgrade": {
                            "type": "boolean"
                        },
                        "AutoMinorVersionUpgrade": {
                            "type": "boolean"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "EngineVersion": {
                            "type": "string"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "MultiAZ": {
                            "type": "boolean"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "PubliclyAccessible": {
                            "type": "boolean"
                        },
                        "ReplicationInstanceClass": {
                            "type": "string"
                        },
                        "ReplicationInstanceIdentifier": {
                            "type": "string"
                        },
                        "ReplicationSubnetGroupIdentifier": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcSecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ReplicationInstanceClass"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DMS::ReplicationInstance"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DMS::ReplicationSubnetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ReplicationSubnetGroupDescription": {
                            "type": "string"
                        },
                        "ReplicationSubnetGroupIdentifier": {
                            "type": "string"
                        },
                        "SubnetIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ReplicationSubnetGroupDescription",
                        "SubnetIds"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DMS::ReplicationSubnetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DMS::ReplicationTask": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CdcStartTime": {
                            "type": "number"
                        },
                        "MigrationType": {
                            "type": "string"
                        },
                        "ReplicationInstanceArn": {
                            "type": "string"
                        },
                        "ReplicationTaskIdentifier": {
                            "type": "string"
                        },
                        "ReplicationTaskSettings": {
                            "type": "string"
                        },
                        "SourceEndpointArn": {
                            "type": "string"
                        },
                        "TableMappings": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TargetEndpointArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "MigrationType",
                        "ReplicationInstanceArn",
                        "SourceEndpointArn",
                        "TableMappings",
                        "TargetEndpointArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DMS::ReplicationTask"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Activate": {
                            "type": "boolean"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "ParameterObjects": {
                            "items": {
                                "$ref": "#/definitions/AWS::DataPipeline::Pipeline.ParameterObject"
                            },
                            "type": "array"
                        },
                        "ParameterValues": {
                            "items": {
                                "$ref": "#/definitions/AWS::DataPipeline::Pipeline.ParameterValue"
                            },
                            "type": "array"
                        },
                        "PipelineObjects": {
                            "items": {
                                "$ref": "#/definitions/AWS::DataPipeline::Pipeline.PipelineObject"
                            },
                            "type": "array"
                        },
                        "PipelineTags": {
                            "items": {
                                "$ref": "#/definitions/AWS::DataPipeline::Pipeline.PipelineTag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name",
                        "ParameterObjects"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DataPipeline::Pipeline"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline.Field": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "RefValue": {
                    "type": "string"
                },
                "StringValue": {
                    "type": "string"
                }
            },
            "required": [
                "Key"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline.ParameterAttribute": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "StringValue": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "StringValue"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline.ParameterObject": {
            "additionalProperties": false,
            "properties": {
                "Attributes": {
                    "items": {
                        "$ref": "#/definitions/AWS::DataPipeline::Pipeline.ParameterAttribute"
                    },
                    "type": "array"
                },
                "Id": {
                    "type": "string"
                }
            },
            "required": [
                "Attributes",
                "Id"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline.ParameterValue": {
            "additionalProperties": false,
            "properties": {
                "Id": {
                    "type": "string"
                },
                "StringValue": {
                    "type": "string"
                }
            },
            "required": [
                "Id",
                "StringValue"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline.PipelineObject": {
            "additionalProperties": false,
            "properties": {
                "Fields": {
                    "items": {
                        "$ref": "#/definitions/AWS::DataPipeline::Pipeline.Field"
                    },
                    "type": "array"
                },
                "Id": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "Fields",
                "Id",
                "Name"
            ],
            "type": "object"
        },
        "AWS::DataPipeline::Pipeline.PipelineTag": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::DirectoryService::MicrosoftAD": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CreateAlias": {
                            "type": "boolean"
                        },
                        "EnableSso": {
                            "type": "boolean"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Password": {
                            "type": "string"
                        },
                        "ShortName": {
                            "type": "string"
                        },
                        "VpcSettings": {
                            "$ref": "#/definitions/AWS::DirectoryService::MicrosoftAD.VpcSettings"
                        }
                    },
                    "required": [
                        "Name",
                        "Password",
                        "VpcSettings"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DirectoryService::MicrosoftAD"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DirectoryService::MicrosoftAD.VpcSettings": {
            "additionalProperties": false,
            "properties": {
                "SubnetIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "VpcId": {
                    "type": "string"
                }
            },
            "required": [
                "SubnetIds",
                "VpcId"
            ],
            "type": "object"
        },
        "AWS::DirectoryService::SimpleAD": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CreateAlias": {
                            "type": "boolean"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "EnableSso": {
                            "type": "boolean"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Password": {
                            "type": "string"
                        },
                        "ShortName": {
                            "type": "string"
                        },
                        "Size": {
                            "type": "string"
                        },
                        "VpcSettings": {
                            "$ref": "#/definitions/AWS::DirectoryService::SimpleAD.VpcSettings"
                        }
                    },
                    "required": [
                        "Name",
                        "Password",
                        "Size",
                        "VpcSettings"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DirectoryService::SimpleAD"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DirectoryService::SimpleAD.VpcSettings": {
            "additionalProperties": false,
            "properties": {
                "SubnetIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "VpcId": {
                    "type": "string"
                }
            },
            "required": [
                "SubnetIds",
                "VpcId"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AttributeDefinitions": {
                            "items": {
                                "$ref": "#/definitions/AWS::DynamoDB::Table.AttributeDefinition"
                            },
                            "type": "array"
                        },
                        "GlobalSecondaryIndexes": {
                            "items": {
                                "$ref": "#/definitions/AWS::DynamoDB::Table.GlobalSecondaryIndex"
                            },
                            "type": "array"
                        },
                        "KeySchema": {
                            "items": {
                                "$ref": "#/definitions/AWS::DynamoDB::Table.KeySchema"
                            },
                            "type": "array"
                        },
                        "LocalSecondaryIndexes": {
                            "items": {
                                "$ref": "#/definitions/AWS::DynamoDB::Table.LocalSecondaryIndex"
                            },
                            "type": "array"
                        },
                        "ProvisionedThroughput": {
                            "$ref": "#/definitions/AWS::DynamoDB::Table.ProvisionedThroughput"
                        },
                        "SSESpecification": {
                            "$ref": "#/definitions/AWS::DynamoDB::Table.SSESpecification"
                        },
                        "StreamSpecification": {
                            "$ref": "#/definitions/AWS::DynamoDB::Table.StreamSpecification"
                        },
                        "TableName": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TimeToLiveSpecification": {
                            "$ref": "#/definitions/AWS::DynamoDB::Table.TimeToLiveSpecification"
                        }
                    },
                    "required": [
                        "KeySchema",
                        "ProvisionedThroughput"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::DynamoDB::Table"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.AttributeDefinition": {
            "additionalProperties": false,
            "properties": {
                "AttributeName": {
                    "type": "string"
                },
                "AttributeType": {
                    "type": "string"
                }
            },
            "required": [
                "AttributeName",
                "AttributeType"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.GlobalSecondaryIndex": {
            "additionalProperties": false,
            "properties": {
                "IndexName": {
                    "type": "string"
                },
                "KeySchema": {
                    "items": {
                        "$ref": "#/definitions/AWS::DynamoDB::Table.KeySchema"
                    },
                    "type": "array"
                },
                "Projection": {
                    "$ref": "#/definitions/AWS::DynamoDB::Table.Projection"
                },
                "ProvisionedThroughput": {
                    "$ref": "#/definitions/AWS::DynamoDB::Table.ProvisionedThroughput"
                }
            },
            "required": [
                "IndexName",
                "KeySchema",
                "Projection",
                "ProvisionedThroughput"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.KeySchema": {
            "additionalProperties": false,
            "properties": {
                "AttributeName": {
                    "type": "string"
                },
                "KeyType": {
                    "type": "string"
                }
            },
            "required": [
                "AttributeName",
                "KeyType"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.LocalSecondaryIndex": {
            "additionalProperties": false,
            "properties": {
                "IndexName": {
                    "type": "string"
                },
                "KeySchema": {
                    "items": {
                        "$ref": "#/definitions/AWS::DynamoDB::Table.KeySchema"
                    },
                    "type": "array"
                },
                "Projection": {
                    "$ref": "#/definitions/AWS::DynamoDB::Table.Projection"
                }
            },
            "required": [
                "IndexName",
                "KeySchema",
                "Projection"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.Projection": {
            "additionalProperties": false,
            "properties": {
                "NonKeyAttributes": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "ProjectionType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::DynamoDB::Table.ProvisionedThroughput": {
            "additionalProperties": false,
            "properties": {
                "ReadCapacityUnits": {
                    "type": "number"
                },
                "WriteCapacityUnits": {
                    "type": "number"
                }
            },
            "required": [
                "ReadCapacityUnits",
                "WriteCapacityUnits"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.SSESpecification": {
            "additionalProperties": false,
            "properties": {
                "SSEEnabled": {
                    "type": "boolean"
                }
            },
            "required": [
                "SSEEnabled"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.StreamSpecification": {
            "additionalProperties": false,
            "properties": {
                "StreamViewType": {
                    "type": "string"
                }
            },
            "required": [
                "StreamViewType"
            ],
            "type": "object"
        },
        "AWS::DynamoDB::Table.TimeToLiveSpecification": {
            "additionalProperties": false,
            "properties": {
                "AttributeName": {
                    "type": "string"
                },
                "Enabled": {
                    "type": "boolean"
                }
            },
            "required": [
                "AttributeName",
                "Enabled"
            ],
            "type": "object"
        },
        "AWS::EC2::CustomerGateway": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "BgpAsn": {
                            "type": "number"
                        },
                        "IpAddress": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "BgpAsn",
                        "IpAddress",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::CustomerGateway"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::DHCPOptions": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DomainName": {
                            "type": "string"
                        },
                        "DomainNameServers": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "NetbiosNameServers": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "NetbiosNodeType": {
                            "type": "number"
                        },
                        "NtpServers": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::DHCPOptions"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EC2::EIP": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Domain": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::EIP"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EC2::EIPAssociation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllocationId": {
                            "type": "string"
                        },
                        "EIP": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "NetworkInterfaceId": {
                            "type": "string"
                        },
                        "PrivateIpAddress": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::EIPAssociation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EC2::EgressOnlyInternetGateway": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::EgressOnlyInternetGateway"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::FlowLog": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DeliverLogsPermissionArn": {
                            "type": "string"
                        },
                        "LogGroupName": {
                            "type": "string"
                        },
                        "ResourceId": {
                            "type": "string"
                        },
                        "ResourceType": {
                            "type": "string"
                        },
                        "TrafficType": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DeliverLogsPermissionArn",
                        "LogGroupName",
                        "ResourceId",
                        "ResourceType",
                        "TrafficType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::FlowLog"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::Host": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutoPlacement": {
                            "type": "string"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "InstanceType": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AvailabilityZone",
                        "InstanceType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::Host"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance": {
            "additionalProperties": false,
            "properties": {
                "CreationPolicy": {
                    "type": "object"
                },
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AdditionalInfo": {
                            "type": "string"
                        },
                        "Affinity": {
                            "type": "string"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "BlockDeviceMappings": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::Instance.BlockDeviceMapping"
                            },
                            "type": "array"
                        },
                        "CreditSpecification": {
                            "$ref": "#/definitions/AWS::EC2::Instance.CreditSpecification"
                        },
                        "DisableApiTermination": {
                            "type": "boolean"
                        },
                        "EbsOptimized": {
                            "type": "boolean"
                        },
                        "ElasticGpuSpecifications": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::Instance.ElasticGpuSpecification"
                            },
                            "type": "array"
                        },
                        "HostId": {
                            "type": "string"
                        },
                        "IamInstanceProfile": {
                            "type": "string"
                        },
                        "ImageId": {
                            "type": "string"
                        },
                        "InstanceInitiatedShutdownBehavior": {
                            "type": "string"
                        },
                        "InstanceType": {
                            "type": "string"
                        },
                        "Ipv6AddressCount": {
                            "type": "number"
                        },
                        "Ipv6Addresses": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::Instance.InstanceIpv6Address"
                            },
                            "type": "array"
                        },
                        "KernelId": {
                            "type": "string"
                        },
                        "KeyName": {
                            "type": "string"
                        },
                        "Monitoring": {
                            "type": "boolean"
                        },
                        "NetworkInterfaces": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::Instance.NetworkInterface"
                            },
                            "type": "array"
                        },
                        "PlacementGroupName": {
                            "type": "string"
                        },
                        "PrivateIpAddress": {
                            "type": "string"
                        },
                        "RamdiskId": {
                            "type": "string"
                        },
                        "SecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SourceDestCheck": {
                            "type": "boolean"
                        },
                        "SsmAssociations": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::Instance.SsmAssociation"
                            },
                            "type": "array"
                        },
                        "SubnetId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Tenancy": {
                            "type": "string"
                        },
                        "UserData": {
                            "type": "string"
                        },
                        "Volumes": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::Instance.Volume"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ImageId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::Instance"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.AssociationParameter": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.BlockDeviceMapping": {
            "additionalProperties": false,
            "properties": {
                "DeviceName": {
                    "type": "string"
                },
                "Ebs": {
                    "$ref": "#/definitions/AWS::EC2::Instance.Ebs"
                },
                "NoDevice": {
                    "$ref": "#/definitions/AWS::EC2::Instance.NoDevice"
                },
                "VirtualName": {
                    "type": "string"
                }
            },
            "required": [
                "DeviceName"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.CreditSpecification": {
            "additionalProperties": false,
            "properties": {
                "CPUCredits": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::Instance.Ebs": {
            "additionalProperties": false,
            "properties": {
                "DeleteOnTermination": {
                    "type": "boolean"
                },
                "Encrypted": {
                    "type": "boolean"
                },
                "Iops": {
                    "type": "number"
                },
                "SnapshotId": {
                    "type": "string"
                },
                "VolumeSize": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::Instance.ElasticGpuSpecification": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.InstanceIpv6Address": {
            "additionalProperties": false,
            "properties": {
                "Ipv6Address": {
                    "type": "string"
                }
            },
            "required": [
                "Ipv6Address"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.NetworkInterface": {
            "additionalProperties": false,
            "properties": {
                "AssociatePublicIpAddress": {
                    "type": "boolean"
                },
                "DeleteOnTermination": {
                    "type": "boolean"
                },
                "Description": {
                    "type": "string"
                },
                "DeviceIndex": {
                    "type": "string"
                },
                "GroupSet": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Ipv6AddressCount": {
                    "type": "number"
                },
                "Ipv6Addresses": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::Instance.InstanceIpv6Address"
                    },
                    "type": "array"
                },
                "NetworkInterfaceId": {
                    "type": "string"
                },
                "PrivateIpAddress": {
                    "type": "string"
                },
                "PrivateIpAddresses": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::Instance.PrivateIpAddressSpecification"
                    },
                    "type": "array"
                },
                "SecondaryPrivateIpAddressCount": {
                    "type": "number"
                },
                "SubnetId": {
                    "type": "string"
                }
            },
            "required": [
                "DeviceIndex"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.NoDevice": {
            "additionalProperties": false,
            "properties": {},
            "type": "object"
        },
        "AWS::EC2::Instance.PrivateIpAddressSpecification": {
            "additionalProperties": false,
            "properties": {
                "Primary": {
                    "type": "boolean"
                },
                "PrivateIpAddress": {
                    "type": "string"
                }
            },
            "required": [
                "Primary",
                "PrivateIpAddress"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.SsmAssociation": {
            "additionalProperties": false,
            "properties": {
                "AssociationParameters": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::Instance.AssociationParameter"
                    },
                    "type": "array"
                },
                "DocumentName": {
                    "type": "string"
                }
            },
            "required": [
                "DocumentName"
            ],
            "type": "object"
        },
        "AWS::EC2::Instance.Volume": {
            "additionalProperties": false,
            "properties": {
                "Device": {
                    "type": "string"
                },
                "VolumeId": {
                    "type": "string"
                }
            },
            "required": [
                "Device",
                "VolumeId"
            ],
            "type": "object"
        },
        "AWS::EC2::InternetGateway": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::InternetGateway"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EC2::NatGateway": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllocationId": {
                            "type": "string"
                        },
                        "SubnetId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "AllocationId",
                        "SubnetId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::NatGateway"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkAcl": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::NetworkAcl"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkAclEntry": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CidrBlock": {
                            "type": "string"
                        },
                        "Egress": {
                            "type": "boolean"
                        },
                        "Icmp": {
                            "$ref": "#/definitions/AWS::EC2::NetworkAclEntry.Icmp"
                        },
                        "Ipv6CidrBlock": {
                            "type": "string"
                        },
                        "NetworkAclId": {
                            "type": "string"
                        },
                        "PortRange": {
                            "$ref": "#/definitions/AWS::EC2::NetworkAclEntry.PortRange"
                        },
                        "Protocol": {
                            "type": "number"
                        },
                        "RuleAction": {
                            "type": "string"
                        },
                        "RuleNumber": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "CidrBlock",
                        "NetworkAclId",
                        "Protocol",
                        "RuleAction",
                        "RuleNumber"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::NetworkAclEntry"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkAclEntry.Icmp": {
            "additionalProperties": false,
            "properties": {
                "Code": {
                    "type": "number"
                },
                "Type": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::EC2::NetworkAclEntry.PortRange": {
            "additionalProperties": false,
            "properties": {
                "From": {
                    "type": "number"
                },
                "To": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::EC2::NetworkInterface": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "GroupSet": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "InterfaceType": {
                            "type": "string"
                        },
                        "Ipv6AddressCount": {
                            "type": "number"
                        },
                        "Ipv6Addresses": {
                            "$ref": "#/definitions/AWS::EC2::NetworkInterface.InstanceIpv6Address"
                        },
                        "PrivateIpAddress": {
                            "type": "string"
                        },
                        "PrivateIpAddresses": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::NetworkInterface.PrivateIpAddressSpecification"
                            },
                            "type": "array"
                        },
                        "SecondaryPrivateIpAddressCount": {
                            "type": "number"
                        },
                        "SourceDestCheck": {
                            "type": "boolean"
                        },
                        "SubnetId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "SubnetId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::NetworkInterface"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkInterface.InstanceIpv6Address": {
            "additionalProperties": false,
            "properties": {
                "Ipv6Address": {
                    "type": "string"
                }
            },
            "required": [
                "Ipv6Address"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkInterface.PrivateIpAddressSpecification": {
            "additionalProperties": false,
            "properties": {
                "Primary": {
                    "type": "boolean"
                },
                "PrivateIpAddress": {
                    "type": "string"
                }
            },
            "required": [
                "Primary",
                "PrivateIpAddress"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkInterfaceAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DeleteOnTermination": {
                            "type": "boolean"
                        },
                        "DeviceIndex": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "NetworkInterfaceId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DeviceIndex",
                        "InstanceId",
                        "NetworkInterfaceId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::NetworkInterfaceAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::NetworkInterfacePermission": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AwsAccountId": {
                            "type": "string"
                        },
                        "NetworkInterfaceId": {
                            "type": "string"
                        },
                        "Permission": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AwsAccountId",
                        "NetworkInterfaceId",
                        "Permission"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::NetworkInterfacePermission"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::PlacementGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Strategy": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::PlacementGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EC2::Route": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DestinationCidrBlock": {
                            "type": "string"
                        },
                        "DestinationIpv6CidrBlock": {
                            "type": "string"
                        },
                        "EgressOnlyInternetGatewayId": {
                            "type": "string"
                        },
                        "GatewayId": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "NatGatewayId": {
                            "type": "string"
                        },
                        "NetworkInterfaceId": {
                            "type": "string"
                        },
                        "RouteTableId": {
                            "type": "string"
                        },
                        "VpcPeeringConnectionId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "RouteTableId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::Route"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::RouteTable": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::RouteTable"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SecurityGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "GroupDescription": {
                            "type": "string"
                        },
                        "GroupName": {
                            "type": "string"
                        },
                        "SecurityGroupEgress": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::SecurityGroup.Egress"
                            },
                            "type": "array"
                        },
                        "SecurityGroupIngress": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::SecurityGroup.Ingress"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "GroupDescription"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SecurityGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SecurityGroup.Egress": {
            "additionalProperties": false,
            "properties": {
                "CidrIp": {
                    "type": "string"
                },
                "CidrIpv6": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "DestinationPrefixListId": {
                    "type": "string"
                },
                "DestinationSecurityGroupId": {
                    "type": "string"
                },
                "FromPort": {
                    "type": "number"
                },
                "IpProtocol": {
                    "type": "string"
                },
                "ToPort": {
                    "type": "number"
                }
            },
            "required": [
                "IpProtocol"
            ],
            "type": "object"
        },
        "AWS::EC2::SecurityGroup.Ingress": {
            "additionalProperties": false,
            "properties": {
                "CidrIp": {
                    "type": "string"
                },
                "CidrIpv6": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "FromPort": {
                    "type": "number"
                },
                "IpProtocol": {
                    "type": "string"
                },
                "SourceSecurityGroupId": {
                    "type": "string"
                },
                "SourceSecurityGroupName": {
                    "type": "string"
                },
                "SourceSecurityGroupOwnerId": {
                    "type": "string"
                },
                "ToPort": {
                    "type": "number"
                }
            },
            "required": [
                "IpProtocol"
            ],
            "type": "object"
        },
        "AWS::EC2::SecurityGroupEgress": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CidrIp": {
                            "type": "string"
                        },
                        "CidrIpv6": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "DestinationPrefixListId": {
                            "type": "string"
                        },
                        "DestinationSecurityGroupId": {
                            "type": "string"
                        },
                        "FromPort": {
                            "type": "number"
                        },
                        "GroupId": {
                            "type": "string"
                        },
                        "IpProtocol": {
                            "type": "string"
                        },
                        "ToPort": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "GroupId",
                        "IpProtocol"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SecurityGroupEgress"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SecurityGroupIngress": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CidrIp": {
                            "type": "string"
                        },
                        "CidrIpv6": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "FromPort": {
                            "type": "number"
                        },
                        "GroupId": {
                            "type": "string"
                        },
                        "GroupName": {
                            "type": "string"
                        },
                        "IpProtocol": {
                            "type": "string"
                        },
                        "SourceSecurityGroupId": {
                            "type": "string"
                        },
                        "SourceSecurityGroupName": {
                            "type": "string"
                        },
                        "SourceSecurityGroupOwnerId": {
                            "type": "string"
                        },
                        "ToPort": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "IpProtocol"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SecurityGroupIngress"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "SpotFleetRequestConfigData": {
                            "$ref": "#/definitions/AWS::EC2::SpotFleet.SpotFleetRequestConfigData"
                        }
                    },
                    "required": [
                        "SpotFleetRequestConfigData"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SpotFleet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.BlockDeviceMapping": {
            "additionalProperties": false,
            "properties": {
                "DeviceName": {
                    "type": "string"
                },
                "Ebs": {
                    "$ref": "#/definitions/AWS::EC2::SpotFleet.EbsBlockDevice"
                },
                "NoDevice": {
                    "type": "string"
                },
                "VirtualName": {
                    "type": "string"
                }
            },
            "required": [
                "DeviceName"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.EbsBlockDevice": {
            "additionalProperties": false,
            "properties": {
                "DeleteOnTermination": {
                    "type": "boolean"
                },
                "Encrypted": {
                    "type": "boolean"
                },
                "Iops": {
                    "type": "number"
                },
                "SnapshotId": {
                    "type": "string"
                },
                "VolumeSize": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.FleetLaunchTemplateSpecification": {
            "additionalProperties": false,
            "properties": {
                "LaunchTemplateId": {
                    "type": "string"
                },
                "LaunchTemplateName": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.GroupIdentifier": {
            "additionalProperties": false,
            "properties": {
                "GroupId": {
                    "type": "string"
                }
            },
            "required": [
                "GroupId"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.IamInstanceProfileSpecification": {
            "additionalProperties": false,
            "properties": {
                "Arn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.InstanceIpv6Address": {
            "additionalProperties": false,
            "properties": {
                "Ipv6Address": {
                    "type": "string"
                }
            },
            "required": [
                "Ipv6Address"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.InstanceNetworkInterfaceSpecification": {
            "additionalProperties": false,
            "properties": {
                "AssociatePublicIpAddress": {
                    "type": "boolean"
                },
                "DeleteOnTermination": {
                    "type": "boolean"
                },
                "Description": {
                    "type": "string"
                },
                "DeviceIndex": {
                    "type": "number"
                },
                "Groups": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Ipv6AddressCount": {
                    "type": "number"
                },
                "Ipv6Addresses": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.InstanceIpv6Address"
                    },
                    "type": "array"
                },
                "NetworkInterfaceId": {
                    "type": "string"
                },
                "PrivateIpAddresses": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.PrivateIpAddressSpecification"
                    },
                    "type": "array"
                },
                "SecondaryPrivateIpAddressCount": {
                    "type": "number"
                },
                "SubnetId": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.LaunchTemplateConfig": {
            "additionalProperties": false,
            "properties": {
                "LaunchTemplateSpecification": {
                    "$ref": "#/definitions/AWS::EC2::SpotFleet.FleetLaunchTemplateSpecification"
                },
                "Overrides": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.LaunchTemplateOverrides"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.LaunchTemplateOverrides": {
            "additionalProperties": false,
            "properties": {
                "AvailabilityZone": {
                    "type": "string"
                },
                "InstanceType": {
                    "type": "string"
                },
                "SpotPrice": {
                    "type": "string"
                },
                "SubnetId": {
                    "type": "string"
                },
                "WeightedCapacity": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.PrivateIpAddressSpecification": {
            "additionalProperties": false,
            "properties": {
                "Primary": {
                    "type": "boolean"
                },
                "PrivateIpAddress": {
                    "type": "string"
                }
            },
            "required": [
                "PrivateIpAddress"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.SpotFleetLaunchSpecification": {
            "additionalProperties": false,
            "properties": {
                "BlockDeviceMappings": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.BlockDeviceMapping"
                    },
                    "type": "array"
                },
                "EbsOptimized": {
                    "type": "boolean"
                },
                "IamInstanceProfile": {
                    "$ref": "#/definitions/AWS::EC2::SpotFleet.IamInstanceProfileSpecification"
                },
                "ImageId": {
                    "type": "string"
                },
                "InstanceType": {
                    "type": "string"
                },
                "KernelId": {
                    "type": "string"
                },
                "KeyName": {
                    "type": "string"
                },
                "Monitoring": {
                    "$ref": "#/definitions/AWS::EC2::SpotFleet.SpotFleetMonitoring"
                },
                "NetworkInterfaces": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.InstanceNetworkInterfaceSpecification"
                    },
                    "type": "array"
                },
                "Placement": {
                    "$ref": "#/definitions/AWS::EC2::SpotFleet.SpotPlacement"
                },
                "RamdiskId": {
                    "type": "string"
                },
                "SecurityGroups": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.GroupIdentifier"
                    },
                    "type": "array"
                },
                "SpotPrice": {
                    "type": "string"
                },
                "SubnetId": {
                    "type": "string"
                },
                "TagSpecifications": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.SpotFleetTagSpecification"
                    },
                    "type": "array"
                },
                "UserData": {
                    "type": "string"
                },
                "WeightedCapacity": {
                    "type": "number"
                }
            },
            "required": [
                "ImageId",
                "InstanceType"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.SpotFleetMonitoring": {
            "additionalProperties": false,
            "properties": {
                "Enabled": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.SpotFleetRequestConfigData": {
            "additionalProperties": false,
            "properties": {
                "AllocationStrategy": {
                    "type": "string"
                },
                "ExcessCapacityTerminationPolicy": {
                    "type": "string"
                },
                "IamFleetRole": {
                    "type": "string"
                },
                "LaunchSpecifications": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.SpotFleetLaunchSpecification"
                    },
                    "type": "array"
                },
                "LaunchTemplateConfigs": {
                    "items": {
                        "$ref": "#/definitions/AWS::EC2::SpotFleet.LaunchTemplateConfig"
                    },
                    "type": "array"
                },
                "ReplaceUnhealthyInstances": {
                    "type": "boolean"
                },
                "SpotPrice": {
                    "type": "string"
                },
                "TargetCapacity": {
                    "type": "number"
                },
                "TerminateInstancesWithExpiration": {
                    "type": "boolean"
                },
                "Type": {
                    "type": "string"
                },
                "ValidFrom": {
                    "type": "string"
                },
                "ValidUntil": {
                    "type": "string"
                }
            },
            "required": [
                "IamFleetRole",
                "TargetCapacity"
            ],
            "type": "object"
        },
        "AWS::EC2::SpotFleet.SpotFleetTagSpecification": {
            "additionalProperties": false,
            "properties": {
                "ResourceType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::SpotFleet.SpotPlacement": {
            "additionalProperties": false,
            "properties": {
                "AvailabilityZone": {
                    "type": "string"
                },
                "GroupName": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::Subnet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AssignIpv6AddressOnCreation": {
                            "type": "boolean"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "CidrBlock": {
                            "type": "string"
                        },
                        "Ipv6CidrBlock": {
                            "type": "string"
                        },
                        "MapPublicIpOnLaunch": {
                            "type": "boolean"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "CidrBlock",
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::Subnet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SubnetCidrBlock": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Ipv6CidrBlock": {
                            "type": "string"
                        },
                        "SubnetId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Ipv6CidrBlock",
                        "SubnetId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SubnetCidrBlock"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SubnetNetworkAclAssociation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "NetworkAclId": {
                            "type": "string"
                        },
                        "SubnetId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "NetworkAclId",
                        "SubnetId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SubnetNetworkAclAssociation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::SubnetRouteTableAssociation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "RouteTableId": {
                            "type": "string"
                        },
                        "SubnetId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "RouteTableId",
                        "SubnetId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::SubnetRouteTableAssociation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::TrunkInterfaceAssociation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "BranchInterfaceId": {
                            "type": "string"
                        },
                        "GREKey": {
                            "type": "number"
                        },
                        "TrunkInterfaceId": {
                            "type": "string"
                        },
                        "VLANId": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "BranchInterfaceId",
                        "TrunkInterfaceId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::TrunkInterfaceAssociation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPC": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CidrBlock": {
                            "type": "string"
                        },
                        "EnableDnsHostnames": {
                            "type": "boolean"
                        },
                        "EnableDnsSupport": {
                            "type": "boolean"
                        },
                        "InstanceTenancy": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "CidrBlock"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPC"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPCCidrBlock": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AmazonProvidedIpv6CidrBlock": {
                            "type": "boolean"
                        },
                        "CidrBlock": {
                            "type": "string"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPCCidrBlock"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPCDHCPOptionsAssociation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DhcpOptionsId": {
                            "type": "string"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DhcpOptionsId",
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPCDHCPOptionsAssociation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPCEndpoint": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PolicyDocument": {
                            "type": "object"
                        },
                        "RouteTableIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ServiceName": {
                            "type": "string"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ServiceName",
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPCEndpoint"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPCGatewayAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "InternetGatewayId": {
                            "type": "string"
                        },
                        "VpcId": {
                            "type": "string"
                        },
                        "VpnGatewayId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPCGatewayAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPCPeeringConnection": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PeerOwnerId": {
                            "type": "string"
                        },
                        "PeerRoleArn": {
                            "type": "string"
                        },
                        "PeerVpcId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "PeerVpcId",
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPCPeeringConnection"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPNConnection": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CustomerGatewayId": {
                            "type": "string"
                        },
                        "StaticRoutesOnly": {
                            "type": "boolean"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Type": {
                            "type": "string"
                        },
                        "VpnGatewayId": {
                            "type": "string"
                        },
                        "VpnTunnelOptionsSpecifications": {
                            "items": {
                                "$ref": "#/definitions/AWS::EC2::VPNConnection.VpnTunnelOptionsSpecification"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "CustomerGatewayId",
                        "Type",
                        "VpnGatewayId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPNConnection"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPNConnection.VpnTunnelOptionsSpecification": {
            "additionalProperties": false,
            "properties": {
                "PreSharedKey": {
                    "type": "string"
                },
                "TunnelInsideCidr": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EC2::VPNConnectionRoute": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DestinationCidrBlock": {
                            "type": "string"
                        },
                        "VpnConnectionId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DestinationCidrBlock",
                        "VpnConnectionId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPNConnectionRoute"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPNGateway": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AmazonSideAsn": {
                            "type": "number"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPNGateway"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VPNGatewayRoutePropagation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "RouteTableIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "VpnGatewayId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "RouteTableIds",
                        "VpnGatewayId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VPNGatewayRoutePropagation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::Volume": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutoEnableIO": {
                            "type": "boolean"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "Encrypted": {
                            "type": "boolean"
                        },
                        "Iops": {
                            "type": "number"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "Size": {
                            "type": "number"
                        },
                        "SnapshotId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VolumeType": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AvailabilityZone"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::Volume"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EC2::VolumeAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Device": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "VolumeId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Device",
                        "InstanceId",
                        "VolumeId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EC2::VolumeAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ECR::Repository": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "LifecyclePolicy": {
                            "$ref": "#/definitions/AWS::ECR::Repository.LifecyclePolicy"
                        },
                        "RepositoryName": {
                            "type": "string"
                        },
                        "RepositoryPolicyText": {
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ECR::Repository"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ECR::Repository.LifecyclePolicy": {
            "additionalProperties": false,
            "properties": {
                "LifecyclePolicyText": {
                    "type": "string"
                },
                "RegistryId": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::Cluster": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ClusterName": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ECS::Cluster"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ECS::Service": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Cluster": {
                            "type": "string"
                        },
                        "DeploymentConfiguration": {
                            "$ref": "#/definitions/AWS::ECS::Service.DeploymentConfiguration"
                        },
                        "DesiredCount": {
                            "type": "number"
                        },
                        "HealthCheckGracePeriodSeconds": {
                            "type": "number"
                        },
                        "LaunchType": {
                            "type": "string"
                        },
                        "LoadBalancers": {
                            "items": {
                                "$ref": "#/definitions/AWS::ECS::Service.LoadBalancer"
                            },
                            "type": "array"
                        },
                        "NetworkConfiguration": {
                            "$ref": "#/definitions/AWS::ECS::Service.NetworkConfiguration"
                        },
                        "PlacementConstraints": {
                            "items": {
                                "$ref": "#/definitions/AWS::ECS::Service.PlacementConstraint"
                            },
                            "type": "array"
                        },
                        "PlacementStrategies": {
                            "items": {
                                "$ref": "#/definitions/AWS::ECS::Service.PlacementStrategy"
                            },
                            "type": "array"
                        },
                        "PlatformVersion": {
                            "type": "string"
                        },
                        "Role": {
                            "type": "string"
                        },
                        "ServiceName": {
                            "type": "string"
                        },
                        "TaskDefinition": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "TaskDefinition"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ECS::Service"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ECS::Service.AwsVpcConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AssignPublicIp": {
                    "type": "string"
                },
                "SecurityGroups": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Subnets": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Subnets"
            ],
            "type": "object"
        },
        "AWS::ECS::Service.DeploymentConfiguration": {
            "additionalProperties": false,
            "properties": {
                "MaximumPercent": {
                    "type": "number"
                },
                "MinimumHealthyPercent": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ECS::Service.LoadBalancer": {
            "additionalProperties": false,
            "properties": {
                "ContainerName": {
                    "type": "string"
                },
                "ContainerPort": {
                    "type": "number"
                },
                "LoadBalancerName": {
                    "type": "string"
                },
                "TargetGroupArn": {
                    "type": "string"
                }
            },
            "required": [
                "ContainerPort"
            ],
            "type": "object"
        },
        "AWS::ECS::Service.NetworkConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AwsvpcConfiguration": {
                    "$ref": "#/definitions/AWS::ECS::Service.AwsVpcConfiguration"
                }
            },
            "type": "object"
        },
        "AWS::ECS::Service.PlacementConstraint": {
            "additionalProperties": false,
            "properties": {
                "Expression": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ECS::Service.PlacementStrategy": {
            "additionalProperties": false,
            "properties": {
                "Field": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ContainerDefinitions": {
                            "items": {
                                "$ref": "#/definitions/AWS::ECS::TaskDefinition.ContainerDefinition"
                            },
                            "type": "array"
                        },
                        "Cpu": {
                            "type": "string"
                        },
                        "ExecutionRoleArn": {
                            "type": "string"
                        },
                        "Family": {
                            "type": "string"
                        },
                        "Memory": {
                            "type": "string"
                        },
                        "NetworkMode": {
                            "type": "string"
                        },
                        "PlacementConstraints": {
                            "items": {
                                "$ref": "#/definitions/AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint"
                            },
                            "type": "array"
                        },
                        "RequiresCompatibilities": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "TaskRoleArn": {
                            "type": "string"
                        },
                        "Volumes": {
                            "items": {
                                "$ref": "#/definitions/AWS::ECS::TaskDefinition.Volume"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ECS::TaskDefinition"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.ContainerDefinition": {
            "additionalProperties": false,
            "properties": {
                "Command": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Cpu": {
                    "type": "number"
                },
                "DisableNetworking": {
                    "type": "boolean"
                },
                "DnsSearchDomains": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "DnsServers": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "DockerLabels": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "DockerSecurityOptions": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "EntryPoint": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Environment": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.KeyValuePair"
                    },
                    "type": "array"
                },
                "Essential": {
                    "type": "boolean"
                },
                "ExtraHosts": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.HostEntry"
                    },
                    "type": "array"
                },
                "Hostname": {
                    "type": "string"
                },
                "Image": {
                    "type": "string"
                },
                "Links": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "LinuxParameters": {
                    "$ref": "#/definitions/AWS::ECS::TaskDefinition.LinuxParameters"
                },
                "LogConfiguration": {
                    "$ref": "#/definitions/AWS::ECS::TaskDefinition.LogConfiguration"
                },
                "Memory": {
                    "type": "number"
                },
                "MemoryReservation": {
                    "type": "number"
                },
                "MountPoints": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.MountPoint"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                },
                "PortMappings": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.PortMapping"
                    },
                    "type": "array"
                },
                "Privileged": {
                    "type": "boolean"
                },
                "ReadonlyRootFilesystem": {
                    "type": "boolean"
                },
                "Ulimits": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.Ulimit"
                    },
                    "type": "array"
                },
                "User": {
                    "type": "string"
                },
                "VolumesFrom": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.VolumeFrom"
                    },
                    "type": "array"
                },
                "WorkingDirectory": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.Device": {
            "additionalProperties": false,
            "properties": {
                "ContainerPath": {
                    "type": "string"
                },
                "HostPath": {
                    "type": "string"
                },
                "Permissions": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "HostPath"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.HostEntry": {
            "additionalProperties": false,
            "properties": {
                "Hostname": {
                    "type": "string"
                },
                "IpAddress": {
                    "type": "string"
                }
            },
            "required": [
                "Hostname",
                "IpAddress"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.HostVolumeProperties": {
            "additionalProperties": false,
            "properties": {
                "SourcePath": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.KernelCapabilities": {
            "additionalProperties": false,
            "properties": {
                "Add": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Drop": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.KeyValuePair": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.LinuxParameters": {
            "additionalProperties": false,
            "properties": {
                "Capabilities": {
                    "$ref": "#/definitions/AWS::ECS::TaskDefinition.KernelCapabilities"
                },
                "Devices": {
                    "items": {
                        "$ref": "#/definitions/AWS::ECS::TaskDefinition.Device"
                    },
                    "type": "array"
                },
                "InitProcessEnabled": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.LogConfiguration": {
            "additionalProperties": false,
            "properties": {
                "LogDriver": {
                    "type": "string"
                },
                "Options": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "required": [
                "LogDriver"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.MountPoint": {
            "additionalProperties": false,
            "properties": {
                "ContainerPath": {
                    "type": "string"
                },
                "ReadOnly": {
                    "type": "boolean"
                },
                "SourceVolume": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.PortMapping": {
            "additionalProperties": false,
            "properties": {
                "ContainerPort": {
                    "type": "number"
                },
                "HostPort": {
                    "type": "number"
                },
                "Protocol": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint": {
            "additionalProperties": false,
            "properties": {
                "Expression": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.Ulimit": {
            "additionalProperties": false,
            "properties": {
                "HardLimit": {
                    "type": "number"
                },
                "Name": {
                    "type": "string"
                },
                "SoftLimit": {
                    "type": "number"
                }
            },
            "required": [
                "HardLimit",
                "Name",
                "SoftLimit"
            ],
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.Volume": {
            "additionalProperties": false,
            "properties": {
                "Host": {
                    "$ref": "#/definitions/AWS::ECS::TaskDefinition.HostVolumeProperties"
                },
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ECS::TaskDefinition.VolumeFrom": {
            "additionalProperties": false,
            "properties": {
                "ReadOnly": {
                    "type": "boolean"
                },
                "SourceContainer": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EFS::FileSystem": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Encrypted": {
                            "type": "boolean"
                        },
                        "FileSystemTags": {
                            "items": {
                                "$ref": "#/definitions/AWS::EFS::FileSystem.ElasticFileSystemTag"
                            },
                            "type": "array"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "PerformanceMode": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EFS::FileSystem"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::EFS::FileSystem.ElasticFileSystemTag": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::EFS::MountTarget": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "FileSystemId": {
                            "type": "string"
                        },
                        "IpAddress": {
                            "type": "string"
                        },
                        "SecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SubnetId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "FileSystemId",
                        "SecurityGroups",
                        "SubnetId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EFS::MountTarget"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AdditionalInfo": {
                            "type": "object"
                        },
                        "Applications": {
                            "items": {
                                "$ref": "#/definitions/AWS::EMR::Cluster.Application"
                            },
                            "type": "array"
                        },
                        "AutoScalingRole": {
                            "type": "string"
                        },
                        "BootstrapActions": {
                            "items": {
                                "$ref": "#/definitions/AWS::EMR::Cluster.BootstrapActionConfig"
                            },
                            "type": "array"
                        },
                        "Configurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::EMR::Cluster.Configuration"
                            },
                            "type": "array"
                        },
                        "CustomAmiId": {
                            "type": "string"
                        },
                        "EbsRootVolumeSize": {
                            "type": "number"
                        },
                        "Instances": {
                            "$ref": "#/definitions/AWS::EMR::Cluster.JobFlowInstancesConfig"
                        },
                        "JobFlowRole": {
                            "type": "string"
                        },
                        "LogUri": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "ReleaseLabel": {
                            "type": "string"
                        },
                        "ScaleDownBehavior": {
                            "type": "string"
                        },
                        "SecurityConfiguration": {
                            "type": "string"
                        },
                        "ServiceRole": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VisibleToAllUsers": {
                            "type": "boolean"
                        }
                    },
                    "required": [
                        "Instances",
                        "JobFlowRole",
                        "Name",
                        "ServiceRole"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EMR::Cluster"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.Application": {
            "additionalProperties": false,
            "properties": {
                "AdditionalInfo": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Args": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::EMR::Cluster.AutoScalingPolicy": {
            "additionalProperties": false,
            "properties": {
                "Constraints": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.ScalingConstraints"
                },
                "Rules": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.ScalingRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Constraints",
                "Rules"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.BootstrapActionConfig": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "ScriptBootstrapAction": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.ScriptBootstrapActionConfig"
                }
            },
            "required": [
                "Name",
                "ScriptBootstrapAction"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.CloudWatchAlarmDefinition": {
            "additionalProperties": false,
            "properties": {
                "ComparisonOperator": {
                    "type": "string"
                },
                "Dimensions": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.MetricDimension"
                    },
                    "type": "array"
                },
                "EvaluationPeriods": {
                    "type": "number"
                },
                "MetricName": {
                    "type": "string"
                },
                "Namespace": {
                    "type": "string"
                },
                "Period": {
                    "type": "number"
                },
                "Statistic": {
                    "type": "string"
                },
                "Threshold": {
                    "type": "number"
                },
                "Unit": {
                    "type": "string"
                }
            },
            "required": [
                "ComparisonOperator",
                "MetricName",
                "Period",
                "Threshold"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.Configuration": {
            "additionalProperties": false,
            "properties": {
                "Classification": {
                    "type": "string"
                },
                "ConfigurationProperties": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Configurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.Configuration"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::EMR::Cluster.EbsBlockDeviceConfig": {
            "additionalProperties": false,
            "properties": {
                "VolumeSpecification": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.VolumeSpecification"
                },
                "VolumesPerInstance": {
                    "type": "number"
                }
            },
            "required": [
                "VolumeSpecification"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.EbsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "EbsBlockDeviceConfigs": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.EbsBlockDeviceConfig"
                    },
                    "type": "array"
                },
                "EbsOptimized": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::EMR::Cluster.InstanceFleetConfig": {
            "additionalProperties": false,
            "properties": {
                "InstanceTypeConfigs": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.InstanceTypeConfig"
                    },
                    "type": "array"
                },
                "LaunchSpecifications": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications"
                },
                "Name": {
                    "type": "string"
                },
                "TargetOnDemandCapacity": {
                    "type": "number"
                },
                "TargetSpotCapacity": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications": {
            "additionalProperties": false,
            "properties": {
                "SpotSpecification": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.SpotProvisioningSpecification"
                }
            },
            "required": [
                "SpotSpecification"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.InstanceGroupConfig": {
            "additionalProperties": false,
            "properties": {
                "AutoScalingPolicy": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.AutoScalingPolicy"
                },
                "BidPrice": {
                    "type": "string"
                },
                "Configurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.Configuration"
                    },
                    "type": "array"
                },
                "EbsConfiguration": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.EbsConfiguration"
                },
                "InstanceCount": {
                    "type": "number"
                },
                "InstanceType": {
                    "type": "string"
                },
                "Market": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "InstanceCount",
                "InstanceType"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.InstanceTypeConfig": {
            "additionalProperties": false,
            "properties": {
                "BidPrice": {
                    "type": "string"
                },
                "BidPriceAsPercentageOfOnDemandPrice": {
                    "type": "number"
                },
                "Configurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Cluster.Configuration"
                    },
                    "type": "array"
                },
                "EbsConfiguration": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.EbsConfiguration"
                },
                "InstanceType": {
                    "type": "string"
                },
                "WeightedCapacity": {
                    "type": "number"
                }
            },
            "required": [
                "InstanceType"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.JobFlowInstancesConfig": {
            "additionalProperties": false,
            "properties": {
                "AdditionalMasterSecurityGroups": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "AdditionalSlaveSecurityGroups": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "CoreInstanceFleet": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.InstanceFleetConfig"
                },
                "CoreInstanceGroup": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.InstanceGroupConfig"
                },
                "Ec2KeyName": {
                    "type": "string"
                },
                "Ec2SubnetId": {
                    "type": "string"
                },
                "EmrManagedMasterSecurityGroup": {
                    "type": "string"
                },
                "EmrManagedSlaveSecurityGroup": {
                    "type": "string"
                },
                "HadoopVersion": {
                    "type": "string"
                },
                "MasterInstanceFleet": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.InstanceFleetConfig"
                },
                "MasterInstanceGroup": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.InstanceGroupConfig"
                },
                "Placement": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.PlacementType"
                },
                "ServiceAccessSecurityGroup": {
                    "type": "string"
                },
                "TerminationProtected": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::EMR::Cluster.MetricDimension": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.PlacementType": {
            "additionalProperties": false,
            "properties": {
                "AvailabilityZone": {
                    "type": "string"
                }
            },
            "required": [
                "AvailabilityZone"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.ScalingAction": {
            "additionalProperties": false,
            "properties": {
                "Market": {
                    "type": "string"
                },
                "SimpleScalingPolicyConfiguration": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.SimpleScalingPolicyConfiguration"
                }
            },
            "required": [
                "SimpleScalingPolicyConfiguration"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.ScalingConstraints": {
            "additionalProperties": false,
            "properties": {
                "MaxCapacity": {
                    "type": "number"
                },
                "MinCapacity": {
                    "type": "number"
                }
            },
            "required": [
                "MaxCapacity",
                "MinCapacity"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.ScalingRule": {
            "additionalProperties": false,
            "properties": {
                "Action": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.ScalingAction"
                },
                "Description": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Trigger": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.ScalingTrigger"
                }
            },
            "required": [
                "Action",
                "Name",
                "Trigger"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.ScalingTrigger": {
            "additionalProperties": false,
            "properties": {
                "CloudWatchAlarmDefinition": {
                    "$ref": "#/definitions/AWS::EMR::Cluster.CloudWatchAlarmDefinition"
                }
            },
            "required": [
                "CloudWatchAlarmDefinition"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.ScriptBootstrapActionConfig": {
            "additionalProperties": false,
            "properties": {
                "Args": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Path": {
                    "type": "string"
                }
            },
            "required": [
                "Path"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.SimpleScalingPolicyConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AdjustmentType": {
                    "type": "string"
                },
                "CoolDown": {
                    "type": "number"
                },
                "ScalingAdjustment": {
                    "type": "number"
                }
            },
            "required": [
                "ScalingAdjustment"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.SpotProvisioningSpecification": {
            "additionalProperties": false,
            "properties": {
                "BlockDurationMinutes": {
                    "type": "number"
                },
                "TimeoutAction": {
                    "type": "string"
                },
                "TimeoutDurationMinutes": {
                    "type": "number"
                }
            },
            "required": [
                "TimeoutAction",
                "TimeoutDurationMinutes"
            ],
            "type": "object"
        },
        "AWS::EMR::Cluster.VolumeSpecification": {
            "additionalProperties": false,
            "properties": {
                "Iops": {
                    "type": "number"
                },
                "SizeInGB": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "required": [
                "SizeInGB",
                "VolumeType"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ClusterId": {
                            "type": "string"
                        },
                        "InstanceFleetType": {
                            "type": "string"
                        },
                        "InstanceTypeConfigs": {
                            "items": {
                                "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.InstanceTypeConfig"
                            },
                            "type": "array"
                        },
                        "LaunchSpecifications": {
                            "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.InstanceFleetProvisioningSpecifications"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "TargetOnDemandCapacity": {
                            "type": "number"
                        },
                        "TargetSpotCapacity": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "ClusterId",
                        "InstanceFleetType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EMR::InstanceFleetConfig"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.Configuration": {
            "additionalProperties": false,
            "properties": {
                "Classification": {
                    "type": "string"
                },
                "ConfigurationProperties": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Configurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.Configuration"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.EbsBlockDeviceConfig": {
            "additionalProperties": false,
            "properties": {
                "VolumeSpecification": {
                    "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.VolumeSpecification"
                },
                "VolumesPerInstance": {
                    "type": "number"
                }
            },
            "required": [
                "VolumeSpecification"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.EbsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "EbsBlockDeviceConfigs": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.EbsBlockDeviceConfig"
                    },
                    "type": "array"
                },
                "EbsOptimized": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.InstanceFleetProvisioningSpecifications": {
            "additionalProperties": false,
            "properties": {
                "SpotSpecification": {
                    "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.SpotProvisioningSpecification"
                }
            },
            "required": [
                "SpotSpecification"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.InstanceTypeConfig": {
            "additionalProperties": false,
            "properties": {
                "BidPrice": {
                    "type": "string"
                },
                "BidPriceAsPercentageOfOnDemandPrice": {
                    "type": "number"
                },
                "Configurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.Configuration"
                    },
                    "type": "array"
                },
                "EbsConfiguration": {
                    "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig.EbsConfiguration"
                },
                "InstanceType": {
                    "type": "string"
                },
                "WeightedCapacity": {
                    "type": "number"
                }
            },
            "required": [
                "InstanceType"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.SpotProvisioningSpecification": {
            "additionalProperties": false,
            "properties": {
                "BlockDurationMinutes": {
                    "type": "number"
                },
                "TimeoutAction": {
                    "type": "string"
                },
                "TimeoutDurationMinutes": {
                    "type": "number"
                }
            },
            "required": [
                "TimeoutAction",
                "TimeoutDurationMinutes"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceFleetConfig.VolumeSpecification": {
            "additionalProperties": false,
            "properties": {
                "Iops": {
                    "type": "number"
                },
                "SizeInGB": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "required": [
                "SizeInGB",
                "VolumeType"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AutoScalingPolicy": {
                            "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.AutoScalingPolicy"
                        },
                        "BidPrice": {
                            "type": "string"
                        },
                        "Configurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.Configuration"
                            },
                            "type": "array"
                        },
                        "EbsConfiguration": {
                            "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.EbsConfiguration"
                        },
                        "InstanceCount": {
                            "type": "number"
                        },
                        "InstanceRole": {
                            "type": "string"
                        },
                        "InstanceType": {
                            "type": "string"
                        },
                        "JobFlowId": {
                            "type": "string"
                        },
                        "Market": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "InstanceCount",
                        "InstanceRole",
                        "InstanceType",
                        "JobFlowId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EMR::InstanceGroupConfig"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.AutoScalingPolicy": {
            "additionalProperties": false,
            "properties": {
                "Constraints": {
                    "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.ScalingConstraints"
                },
                "Rules": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.ScalingRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Constraints",
                "Rules"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition": {
            "additionalProperties": false,
            "properties": {
                "ComparisonOperator": {
                    "type": "string"
                },
                "Dimensions": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.MetricDimension"
                    },
                    "type": "array"
                },
                "EvaluationPeriods": {
                    "type": "number"
                },
                "MetricName": {
                    "type": "string"
                },
                "Namespace": {
                    "type": "string"
                },
                "Period": {
                    "type": "number"
                },
                "Statistic": {
                    "type": "string"
                },
                "Threshold": {
                    "type": "number"
                },
                "Unit": {
                    "type": "string"
                }
            },
            "required": [
                "ComparisonOperator",
                "MetricName",
                "Period",
                "Threshold"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.Configuration": {
            "additionalProperties": false,
            "properties": {
                "Classification": {
                    "type": "string"
                },
                "ConfigurationProperties": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Configurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.Configuration"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig": {
            "additionalProperties": false,
            "properties": {
                "VolumeSpecification": {
                    "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.VolumeSpecification"
                },
                "VolumesPerInstance": {
                    "type": "number"
                }
            },
            "required": [
                "VolumeSpecification"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.EbsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "EbsBlockDeviceConfigs": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig"
                    },
                    "type": "array"
                },
                "EbsOptimized": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.MetricDimension": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.ScalingAction": {
            "additionalProperties": false,
            "properties": {
                "Market": {
                    "type": "string"
                },
                "SimpleScalingPolicyConfiguration": {
                    "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration"
                }
            },
            "required": [
                "SimpleScalingPolicyConfiguration"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.ScalingConstraints": {
            "additionalProperties": false,
            "properties": {
                "MaxCapacity": {
                    "type": "number"
                },
                "MinCapacity": {
                    "type": "number"
                }
            },
            "required": [
                "MaxCapacity",
                "MinCapacity"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.ScalingRule": {
            "additionalProperties": false,
            "properties": {
                "Action": {
                    "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.ScalingAction"
                },
                "Description": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Trigger": {
                    "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.ScalingTrigger"
                }
            },
            "required": [
                "Action",
                "Name",
                "Trigger"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.ScalingTrigger": {
            "additionalProperties": false,
            "properties": {
                "CloudWatchAlarmDefinition": {
                    "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition"
                }
            },
            "required": [
                "CloudWatchAlarmDefinition"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AdjustmentType": {
                    "type": "string"
                },
                "CoolDown": {
                    "type": "number"
                },
                "ScalingAdjustment": {
                    "type": "number"
                }
            },
            "required": [
                "ScalingAdjustment"
            ],
            "type": "object"
        },
        "AWS::EMR::InstanceGroupConfig.VolumeSpecification": {
            "additionalProperties": false,
            "properties": {
                "Iops": {
                    "type": "number"
                },
                "SizeInGB": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "required": [
                "SizeInGB",
                "VolumeType"
            ],
            "type": "object"
        },
        "AWS::EMR::SecurityConfiguration": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "SecurityConfiguration": {
                            "type": "object"
                        }
                    },
                    "required": [
                        "SecurityConfiguration"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EMR::SecurityConfiguration"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EMR::Step": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ActionOnFailure": {
                            "type": "string"
                        },
                        "HadoopJarStep": {
                            "$ref": "#/definitions/AWS::EMR::Step.HadoopJarStepConfig"
                        },
                        "JobFlowId": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ActionOnFailure",
                        "HadoopJarStep",
                        "JobFlowId",
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::EMR::Step"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::EMR::Step.HadoopJarStepConfig": {
            "additionalProperties": false,
            "properties": {
                "Args": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Jar": {
                    "type": "string"
                },
                "MainClass": {
                    "type": "string"
                },
                "StepProperties": {
                    "items": {
                        "$ref": "#/definitions/AWS::EMR::Step.KeyValue"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Jar"
            ],
            "type": "object"
        },
        "AWS::EMR::Step.KeyValue": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElastiCache::CacheCluster": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AZMode": {
                            "type": "string"
                        },
                        "AutoMinorVersionUpgrade": {
                            "type": "boolean"
                        },
                        "CacheNodeType": {
                            "type": "string"
                        },
                        "CacheParameterGroupName": {
                            "type": "string"
                        },
                        "CacheSecurityGroupNames": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "CacheSubnetGroupName": {
                            "type": "string"
                        },
                        "ClusterName": {
                            "type": "string"
                        },
                        "Engine": {
                            "type": "string"
                        },
                        "EngineVersion": {
                            "type": "string"
                        },
                        "NotificationTopicArn": {
                            "type": "string"
                        },
                        "NumCacheNodes": {
                            "type": "number"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "PreferredAvailabilityZone": {
                            "type": "string"
                        },
                        "PreferredAvailabilityZones": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "SnapshotArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SnapshotName": {
                            "type": "string"
                        },
                        "SnapshotRetentionLimit": {
                            "type": "number"
                        },
                        "SnapshotWindow": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcSecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "CacheNodeType",
                        "Engine",
                        "NumCacheNodes"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElastiCache::CacheCluster"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElastiCache::ParameterGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CacheParameterGroupFamily": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Properties": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "required": [
                        "CacheParameterGroupFamily",
                        "Description"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElastiCache::ParameterGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElastiCache::ReplicationGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AtRestEncryptionEnabled": {
                            "type": "boolean"
                        },
                        "AuthToken": {
                            "type": "string"
                        },
                        "AutoMinorVersionUpgrade": {
                            "type": "boolean"
                        },
                        "AutomaticFailoverEnabled": {
                            "type": "boolean"
                        },
                        "CacheNodeType": {
                            "type": "string"
                        },
                        "CacheParameterGroupName": {
                            "type": "string"
                        },
                        "CacheSecurityGroupNames": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "CacheSubnetGroupName": {
                            "type": "string"
                        },
                        "Engine": {
                            "type": "string"
                        },
                        "EngineVersion": {
                            "type": "string"
                        },
                        "NodeGroupConfiguration": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration"
                            },
                            "type": "array"
                        },
                        "NotificationTopicArn": {
                            "type": "string"
                        },
                        "NumCacheClusters": {
                            "type": "number"
                        },
                        "NumNodeGroups": {
                            "type": "number"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "PreferredCacheClusterAZs": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "PrimaryClusterId": {
                            "type": "string"
                        },
                        "ReplicasPerNodeGroup": {
                            "type": "number"
                        },
                        "ReplicationGroupDescription": {
                            "type": "string"
                        },
                        "ReplicationGroupId": {
                            "type": "string"
                        },
                        "SecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SnapshotArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SnapshotName": {
                            "type": "string"
                        },
                        "SnapshotRetentionLimit": {
                            "type": "number"
                        },
                        "SnapshotWindow": {
                            "type": "string"
                        },
                        "SnapshottingClusterId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TransitEncryptionEnabled": {
                            "type": "boolean"
                        }
                    },
                    "required": [
                        "ReplicationGroupDescription"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElastiCache::ReplicationGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration": {
            "additionalProperties": false,
            "properties": {
                "PrimaryAvailabilityZone": {
                    "type": "string"
                },
                "ReplicaAvailabilityZones": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "ReplicaCount": {
                    "type": "number"
                },
                "Slots": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElastiCache::SecurityGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Description"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElastiCache::SecurityGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElastiCache::SecurityGroupIngress": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CacheSecurityGroupName": {
                            "type": "string"
                        },
                        "EC2SecurityGroupName": {
                            "type": "string"
                        },
                        "EC2SecurityGroupOwnerId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "CacheSecurityGroupName",
                        "EC2SecurityGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElastiCache::SecurityGroupIngress"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElastiCache::SubnetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CacheSubnetGroupName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "SubnetIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Description",
                        "SubnetIds"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElastiCache::SubnetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Application": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "ResourceLifecycleConfig": {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::Application.ApplicationResourceLifecycleConfig"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticBeanstalk::Application"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Application.ApplicationResourceLifecycleConfig": {
            "additionalProperties": false,
            "properties": {
                "ServiceRole": {
                    "type": "string"
                },
                "VersionLifecycleConfig": {
                    "$ref": "#/definitions/AWS::ElasticBeanstalk::Application.ApplicationVersionLifecycleConfig"
                }
            },
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Application.ApplicationVersionLifecycleConfig": {
            "additionalProperties": false,
            "properties": {
                "MaxAgeRule": {
                    "$ref": "#/definitions/AWS::ElasticBeanstalk::Application.MaxAgeRule"
                },
                "MaxCountRule": {
                    "$ref": "#/definitions/AWS::ElasticBeanstalk::Application.MaxCountRule"
                }
            },
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Application.MaxAgeRule": {
            "additionalProperties": false,
            "properties": {
                "DeleteSourceFromS3": {
                    "type": "boolean"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "MaxAgeInDays": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Application.MaxCountRule": {
            "additionalProperties": false,
            "properties": {
                "DeleteSourceFromS3": {
                    "type": "boolean"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "MaxCount": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::ElasticBeanstalk::ApplicationVersion": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "SourceBundle": {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::ApplicationVersion.SourceBundle"
                        }
                    },
                    "required": [
                        "ApplicationName",
                        "SourceBundle"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticBeanstalk::ApplicationVersion"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::ApplicationVersion.SourceBundle": {
            "additionalProperties": false,
            "properties": {
                "S3Bucket": {
                    "type": "string"
                },
                "S3Key": {
                    "type": "string"
                }
            },
            "required": [
                "S3Bucket",
                "S3Key"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::ConfigurationTemplate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "EnvironmentId": {
                            "type": "string"
                        },
                        "OptionSettings": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting"
                            },
                            "type": "array"
                        },
                        "PlatformArn": {
                            "type": "string"
                        },
                        "SolutionStackName": {
                            "type": "string"
                        },
                        "SourceConfiguration": {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration"
                        }
                    },
                    "required": [
                        "ApplicationName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticBeanstalk::ConfigurationTemplate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting": {
            "additionalProperties": false,
            "properties": {
                "Namespace": {
                    "type": "string"
                },
                "OptionName": {
                    "type": "string"
                },
                "ResourceName": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Namespace",
                "OptionName"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ApplicationName": {
                    "type": "string"
                },
                "TemplateName": {
                    "type": "string"
                }
            },
            "required": [
                "ApplicationName",
                "TemplateName"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Environment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "CNAMEPrefix": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "EnvironmentName": {
                            "type": "string"
                        },
                        "OptionSettings": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticBeanstalk::Environment.OptionSetting"
                            },
                            "type": "array"
                        },
                        "PlatformArn": {
                            "type": "string"
                        },
                        "SolutionStackName": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TemplateName": {
                            "type": "string"
                        },
                        "Tier": {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::Environment.Tier"
                        },
                        "VersionLabel": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ApplicationName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticBeanstalk::Environment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Environment.OptionSetting": {
            "additionalProperties": false,
            "properties": {
                "Namespace": {
                    "type": "string"
                },
                "OptionName": {
                    "type": "string"
                },
                "ResourceName": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Namespace",
                "OptionName"
            ],
            "type": "object"
        },
        "AWS::ElasticBeanstalk::Environment.Tier": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AccessLoggingPolicy": {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.AccessLoggingPolicy"
                        },
                        "AppCookieStickinessPolicy": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy"
                            },
                            "type": "array"
                        },
                        "AvailabilityZones": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ConnectionDrainingPolicy": {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy"
                        },
                        "ConnectionSettings": {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings"
                        },
                        "CrossZone": {
                            "type": "boolean"
                        },
                        "HealthCheck": {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck"
                        },
                        "Instances": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "LBCookieStickinessPolicy": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.LBCookieStickinessPolicy"
                            },
                            "type": "array"
                        },
                        "Listeners": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.Listeners"
                            },
                            "type": "array"
                        },
                        "LoadBalancerName": {
                            "type": "string"
                        },
                        "Policies": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer.Policies"
                            },
                            "type": "array"
                        },
                        "Scheme": {
                            "type": "string"
                        },
                        "SecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Subnets": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Listeners"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticLoadBalancing::LoadBalancer"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.AccessLoggingPolicy": {
            "additionalProperties": false,
            "properties": {
                "EmitInterval": {
                    "type": "number"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "S3BucketName": {
                    "type": "string"
                },
                "S3BucketPrefix": {
                    "type": "string"
                }
            },
            "required": [
                "Enabled",
                "S3BucketName"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy": {
            "additionalProperties": false,
            "properties": {
                "CookieName": {
                    "type": "string"
                },
                "PolicyName": {
                    "type": "string"
                }
            },
            "required": [
                "CookieName",
                "PolicyName"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy": {
            "additionalProperties": false,
            "properties": {
                "Enabled": {
                    "type": "boolean"
                },
                "Timeout": {
                    "type": "number"
                }
            },
            "required": [
                "Enabled"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings": {
            "additionalProperties": false,
            "properties": {
                "IdleTimeout": {
                    "type": "number"
                }
            },
            "required": [
                "IdleTimeout"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck": {
            "additionalProperties": false,
            "properties": {
                "HealthyThreshold": {
                    "type": "string"
                },
                "Interval": {
                    "type": "string"
                },
                "Target": {
                    "type": "string"
                },
                "Timeout": {
                    "type": "string"
                },
                "UnhealthyThreshold": {
                    "type": "string"
                }
            },
            "required": [
                "HealthyThreshold",
                "Interval",
                "Target",
                "Timeout",
                "UnhealthyThreshold"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.LBCookieStickinessPolicy": {
            "additionalProperties": false,
            "properties": {
                "CookieExpirationPeriod": {
                    "type": "string"
                },
                "PolicyName": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.Listeners": {
            "additionalProperties": false,
            "properties": {
                "InstancePort": {
                    "type": "string"
                },
                "InstanceProtocol": {
                    "type": "string"
                },
                "LoadBalancerPort": {
                    "type": "string"
                },
                "PolicyNames": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Protocol": {
                    "type": "string"
                },
                "SSLCertificateId": {
                    "type": "string"
                }
            },
            "required": [
                "InstancePort",
                "LoadBalancerPort",
                "Protocol"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancing::LoadBalancer.Policies": {
            "additionalProperties": false,
            "properties": {
                "Attributes": {
                    "items": {
                        "type": "object"
                    },
                    "type": "array"
                },
                "InstancePorts": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "LoadBalancerPorts": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "PolicyName": {
                    "type": "string"
                },
                "PolicyType": {
                    "type": "string"
                }
            },
            "required": [
                "Attributes",
                "PolicyName",
                "PolicyType"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::Listener": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Certificates": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::Listener.Certificate"
                            },
                            "type": "array"
                        },
                        "DefaultActions": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::Listener.Action"
                            },
                            "type": "array"
                        },
                        "LoadBalancerArn": {
                            "type": "string"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "Protocol": {
                            "type": "string"
                        },
                        "SslPolicy": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DefaultActions",
                        "LoadBalancerArn",
                        "Port",
                        "Protocol"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticLoadBalancingV2::Listener"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::Listener.Action": {
            "additionalProperties": false,
            "properties": {
                "TargetGroupArn": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "TargetGroupArn",
                "Type"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::Listener.Certificate": {
            "additionalProperties": false,
            "properties": {
                "CertificateArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::ListenerCertificate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Certificates": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::ListenerCertificate.Certificate"
                            },
                            "type": "array"
                        },
                        "ListenerArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Certificates",
                        "ListenerArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticLoadBalancingV2::ListenerCertificate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::ListenerCertificate.Certificate": {
            "additionalProperties": false,
            "properties": {
                "CertificateArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::ListenerRule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Actions": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::ListenerRule.Action"
                            },
                            "type": "array"
                        },
                        "Conditions": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::ListenerRule.RuleCondition"
                            },
                            "type": "array"
                        },
                        "ListenerArn": {
                            "type": "string"
                        },
                        "Priority": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "Actions",
                        "Conditions",
                        "ListenerArn",
                        "Priority"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticLoadBalancingV2::ListenerRule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::ListenerRule.Action": {
            "additionalProperties": false,
            "properties": {
                "TargetGroupArn": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "TargetGroupArn",
                "Type"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::ListenerRule.RuleCondition": {
            "additionalProperties": false,
            "properties": {
                "Field": {
                    "type": "string"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::LoadBalancer": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "IpAddressType": {
                            "type": "string"
                        },
                        "LoadBalancerAttributes": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Scheme": {
                            "type": "string"
                        },
                        "SecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SubnetMappings": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::LoadBalancer.SubnetMapping"
                            },
                            "type": "array"
                        },
                        "Subnets": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticLoadBalancingV2::LoadBalancer"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::LoadBalancer.SubnetMapping": {
            "additionalProperties": false,
            "properties": {
                "AllocationId": {
                    "type": "string"
                },
                "SubnetId": {
                    "type": "string"
                }
            },
            "required": [
                "AllocationId",
                "SubnetId"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::TargetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "HealthCheckIntervalSeconds": {
                            "type": "number"
                        },
                        "HealthCheckPath": {
                            "type": "string"
                        },
                        "HealthCheckPort": {
                            "type": "string"
                        },
                        "HealthCheckProtocol": {
                            "type": "string"
                        },
                        "HealthCheckTimeoutSeconds": {
                            "type": "number"
                        },
                        "HealthyThresholdCount": {
                            "type": "number"
                        },
                        "Matcher": {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::TargetGroup.Matcher"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "Protocol": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "TargetGroupAttributes": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute"
                            },
                            "type": "array"
                        },
                        "TargetType": {
                            "type": "string"
                        },
                        "Targets": {
                            "items": {
                                "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription"
                            },
                            "type": "array"
                        },
                        "UnhealthyThresholdCount": {
                            "type": "number"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Port",
                        "Protocol",
                        "VpcId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ElasticLoadBalancingV2::TargetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::TargetGroup.Matcher": {
            "additionalProperties": false,
            "properties": {
                "HttpCode": {
                    "type": "string"
                }
            },
            "required": [
                "HttpCode"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription": {
            "additionalProperties": false,
            "properties": {
                "AvailabilityZone": {
                    "type": "string"
                },
                "Id": {
                    "type": "string"
                },
                "Port": {
                    "type": "number"
                }
            },
            "required": [
                "Id"
            ],
            "type": "object"
        },
        "AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Elasticsearch::Domain": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AccessPolicies": {
                            "type": "object"
                        },
                        "AdvancedOptions": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "DomainName": {
                            "type": "string"
                        },
                        "EBSOptions": {
                            "$ref": "#/definitions/AWS::Elasticsearch::Domain.EBSOptions"
                        },
                        "ElasticsearchClusterConfig": {
                            "$ref": "#/definitions/AWS::Elasticsearch::Domain.ElasticsearchClusterConfig"
                        },
                        "ElasticsearchVersion": {
                            "type": "string"
                        },
                        "SnapshotOptions": {
                            "$ref": "#/definitions/AWS::Elasticsearch::Domain.SnapshotOptions"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VPCOptions": {
                            "$ref": "#/definitions/AWS::Elasticsearch::Domain.VPCOptions"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Elasticsearch::Domain"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Elasticsearch::Domain.EBSOptions": {
            "additionalProperties": false,
            "properties": {
                "EBSEnabled": {
                    "type": "boolean"
                },
                "Iops": {
                    "type": "number"
                },
                "VolumeSize": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Elasticsearch::Domain.ElasticsearchClusterConfig": {
            "additionalProperties": false,
            "properties": {
                "DedicatedMasterCount": {
                    "type": "number"
                },
                "DedicatedMasterEnabled": {
                    "type": "boolean"
                },
                "DedicatedMasterType": {
                    "type": "string"
                },
                "InstanceCount": {
                    "type": "number"
                },
                "InstanceType": {
                    "type": "string"
                },
                "ZoneAwarenessEnabled": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::Elasticsearch::Domain.SnapshotOptions": {
            "additionalProperties": false,
            "properties": {
                "AutomatedSnapshotStartHour": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::Elasticsearch::Domain.VPCOptions": {
            "additionalProperties": false,
            "properties": {
                "SecurityGroupIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SubnetIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::Events::Rule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "EventPattern": {
                            "type": "object"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "RoleArn": {
                            "type": "string"
                        },
                        "ScheduleExpression": {
                            "type": "string"
                        },
                        "State": {
                            "type": "string"
                        },
                        "Targets": {
                            "items": {
                                "$ref": "#/definitions/AWS::Events::Rule.Target"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Events::Rule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Events::Rule.EcsParameters": {
            "additionalProperties": false,
            "properties": {
                "TaskCount": {
                    "type": "number"
                },
                "TaskDefinitionArn": {
                    "type": "string"
                }
            },
            "required": [
                "TaskDefinitionArn"
            ],
            "type": "object"
        },
        "AWS::Events::Rule.InputTransformer": {
            "additionalProperties": false,
            "properties": {
                "InputPathsMap": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "InputTemplate": {
                    "type": "string"
                }
            },
            "required": [
                "InputTemplate"
            ],
            "type": "object"
        },
        "AWS::Events::Rule.KinesisParameters": {
            "additionalProperties": false,
            "properties": {
                "PartitionKeyPath": {
                    "type": "string"
                }
            },
            "required": [
                "PartitionKeyPath"
            ],
            "type": "object"
        },
        "AWS::Events::Rule.RunCommandParameters": {
            "additionalProperties": false,
            "properties": {
                "RunCommandTargets": {
                    "items": {
                        "$ref": "#/definitions/AWS::Events::Rule.RunCommandTarget"
                    },
                    "type": "array"
                }
            },
            "required": [
                "RunCommandTargets"
            ],
            "type": "object"
        },
        "AWS::Events::Rule.RunCommandTarget": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Key",
                "Values"
            ],
            "type": "object"
        },
        "AWS::Events::Rule.Target": {
            "additionalProperties": false,
            "properties": {
                "Arn": {
                    "type": "string"
                },
                "EcsParameters": {
                    "$ref": "#/definitions/AWS::Events::Rule.EcsParameters"
                },
                "Id": {
                    "type": "string"
                },
                "Input": {
                    "type": "string"
                },
                "InputPath": {
                    "type": "string"
                },
                "InputTransformer": {
                    "$ref": "#/definitions/AWS::Events::Rule.InputTransformer"
                },
                "KinesisParameters": {
                    "$ref": "#/definitions/AWS::Events::Rule.KinesisParameters"
                },
                "RoleArn": {
                    "type": "string"
                },
                "RunCommandParameters": {
                    "$ref": "#/definitions/AWS::Events::Rule.RunCommandParameters"
                }
            },
            "required": [
                "Arn",
                "Id"
            ],
            "type": "object"
        },
        "AWS::GameLift::Alias": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "RoutingStrategy": {
                            "$ref": "#/definitions/AWS::GameLift::Alias.RoutingStrategy"
                        }
                    },
                    "required": [
                        "Name",
                        "RoutingStrategy"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GameLift::Alias"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::GameLift::Alias.RoutingStrategy": {
            "additionalProperties": false,
            "properties": {
                "FleetId": {
                    "type": "string"
                },
                "Message": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::GameLift::Build": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "StorageLocation": {
                            "$ref": "#/definitions/AWS::GameLift::Build.S3Location"
                        },
                        "Version": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GameLift::Build"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::GameLift::Build.S3Location": {
            "additionalProperties": false,
            "properties": {
                "Bucket": {
                    "type": "string"
                },
                "Key": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                }
            },
            "required": [
                "Bucket",
                "Key",
                "RoleArn"
            ],
            "type": "object"
        },
        "AWS::GameLift::Fleet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "BuildId": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "DesiredEC2Instances": {
                            "type": "number"
                        },
                        "EC2InboundPermissions": {
                            "items": {
                                "$ref": "#/definitions/AWS::GameLift::Fleet.IpPermission"
                            },
                            "type": "array"
                        },
                        "EC2InstanceType": {
                            "type": "string"
                        },
                        "LogPaths": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "MaxSize": {
                            "type": "number"
                        },
                        "MinSize": {
                            "type": "number"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "ServerLaunchParameters": {
                            "type": "string"
                        },
                        "ServerLaunchPath": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "BuildId",
                        "DesiredEC2Instances",
                        "EC2InstanceType",
                        "Name",
                        "ServerLaunchPath"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GameLift::Fleet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::GameLift::Fleet.IpPermission": {
            "additionalProperties": false,
            "properties": {
                "FromPort": {
                    "type": "number"
                },
                "IpRange": {
                    "type": "string"
                },
                "Protocol": {
                    "type": "string"
                },
                "ToPort": {
                    "type": "number"
                }
            },
            "required": [
                "FromPort",
                "IpRange",
                "Protocol",
                "ToPort"
            ],
            "type": "object"
        },
        "AWS::Glue::Classifier": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "GrokClassifier": {
                            "$ref": "#/definitions/AWS::Glue::Classifier.GrokClassifier"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Classifier"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Glue::Classifier.GrokClassifier": {
            "additionalProperties": false,
            "properties": {
                "Classification": {
                    "type": "string"
                },
                "CustomPatterns": {
                    "type": "string"
                },
                "GrokPattern": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "Classification",
                "GrokPattern"
            ],
            "type": "object"
        },
        "AWS::Glue::Connection": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CatalogId": {
                            "type": "string"
                        },
                        "ConnectionInput": {
                            "$ref": "#/definitions/AWS::Glue::Connection.ConnectionInput"
                        }
                    },
                    "required": [
                        "CatalogId",
                        "ConnectionInput"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Connection"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Connection.ConnectionInput": {
            "additionalProperties": false,
            "properties": {
                "ConnectionProperties": {
                    "type": "object"
                },
                "ConnectionType": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "MatchCriteria": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                },
                "PhysicalConnectionRequirements": {
                    "$ref": "#/definitions/AWS::Glue::Connection.PhysicalConnectionRequirements"
                }
            },
            "required": [
                "ConnectionProperties",
                "ConnectionType"
            ],
            "type": "object"
        },
        "AWS::Glue::Connection.PhysicalConnectionRequirements": {
            "additionalProperties": false,
            "properties": {
                "AvailabilityZone": {
                    "type": "string"
                },
                "SecurityGroupIdList": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SubnetId": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Crawler": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Classifiers": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "DatabaseName": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Role": {
                            "type": "string"
                        },
                        "Schedule": {
                            "$ref": "#/definitions/AWS::Glue::Crawler.Schedule"
                        },
                        "SchemaChangePolicy": {
                            "$ref": "#/definitions/AWS::Glue::Crawler.SchemaChangePolicy"
                        },
                        "TablePrefix": {
                            "type": "string"
                        },
                        "Targets": {
                            "$ref": "#/definitions/AWS::Glue::Crawler.Targets"
                        }
                    },
                    "required": [
                        "DatabaseName",
                        "Role",
                        "Targets"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Crawler"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Crawler.JdbcTarget": {
            "additionalProperties": false,
            "properties": {
                "ConnectionName": {
                    "type": "string"
                },
                "Exclusions": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Path": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Crawler.S3Target": {
            "additionalProperties": false,
            "properties": {
                "Exclusions": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Path": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Crawler.Schedule": {
            "additionalProperties": false,
            "properties": {
                "ScheduleExpression": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Crawler.SchemaChangePolicy": {
            "additionalProperties": false,
            "properties": {
                "DeleteBehavior": {
                    "type": "string"
                },
                "UpdateBehavior": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Crawler.Targets": {
            "additionalProperties": false,
            "properties": {
                "JdbcTargets": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Crawler.JdbcTarget"
                    },
                    "type": "array"
                },
                "S3Targets": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Crawler.S3Target"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Database": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CatalogId": {
                            "type": "string"
                        },
                        "DatabaseInput": {
                            "$ref": "#/definitions/AWS::Glue::Database.DatabaseInput"
                        }
                    },
                    "required": [
                        "CatalogId",
                        "DatabaseInput"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Database"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Database.DatabaseInput": {
            "additionalProperties": false,
            "properties": {
                "Description": {
                    "type": "string"
                },
                "LocationUri": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                }
            },
            "type": "object"
        },
        "AWS::Glue::DevEndpoint": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "EndpointName": {
                            "type": "string"
                        },
                        "ExtraJarsS3Path": {
                            "type": "string"
                        },
                        "ExtraPythonLibsS3Path": {
                            "type": "string"
                        },
                        "NumberOfNodes": {
                            "type": "number"
                        },
                        "PublicKey": {
                            "type": "string"
                        },
                        "RoleArn": {
                            "type": "string"
                        },
                        "SecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SubnetId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "PublicKey",
                        "RoleArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::DevEndpoint"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Job": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllocatedCapacity": {
                            "type": "number"
                        },
                        "Command": {
                            "$ref": "#/definitions/AWS::Glue::Job.JobCommand"
                        },
                        "Connections": {
                            "$ref": "#/definitions/AWS::Glue::Job.ConnectionsList"
                        },
                        "DefaultArguments": {
                            "type": "object"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "ExecutionProperty": {
                            "$ref": "#/definitions/AWS::Glue::Job.ExecutionProperty"
                        },
                        "LogUri": {
                            "type": "string"
                        },
                        "MaxRetries": {
                            "type": "number"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Role": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Command",
                        "Role"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Job"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Job.ConnectionsList": {
            "additionalProperties": false,
            "properties": {
                "Connections": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Job.ExecutionProperty": {
            "additionalProperties": false,
            "properties": {
                "MaxConcurrentRuns": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Job.JobCommand": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "ScriptLocation": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Partition": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CatalogId": {
                            "type": "string"
                        },
                        "DatabaseName": {
                            "type": "string"
                        },
                        "PartitionInput": {
                            "$ref": "#/definitions/AWS::Glue::Partition.PartitionInput"
                        },
                        "TableName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "CatalogId",
                        "DatabaseName",
                        "PartitionInput",
                        "TableName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Partition"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Partition.Column": {
            "additionalProperties": false,
            "properties": {
                "Comment": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Name"
            ],
            "type": "object"
        },
        "AWS::Glue::Partition.Order": {
            "additionalProperties": false,
            "properties": {
                "Column": {
                    "type": "string"
                },
                "SortOrder": {
                    "type": "number"
                }
            },
            "required": [
                "Column"
            ],
            "type": "object"
        },
        "AWS::Glue::Partition.PartitionInput": {
            "additionalProperties": false,
            "properties": {
                "Parameters": {
                    "type": "object"
                },
                "StorageDescriptor": {
                    "$ref": "#/definitions/AWS::Glue::Partition.StorageDescriptor"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Values"
            ],
            "type": "object"
        },
        "AWS::Glue::Partition.SerdeInfo": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                },
                "SerializationLibrary": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Partition.SkewedInfo": {
            "additionalProperties": false,
            "properties": {
                "SkewedColumnNames": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SkewedColumnValueLocationMaps": {
                    "type": "object"
                },
                "SkewedColumnValues": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Partition.StorageDescriptor": {
            "additionalProperties": false,
            "properties": {
                "BucketColumns": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Columns": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Partition.Column"
                    },
                    "type": "array"
                },
                "Compressed": {
                    "type": "boolean"
                },
                "InputFormat": {
                    "type": "string"
                },
                "Location": {
                    "type": "string"
                },
                "NumberOfBuckets": {
                    "type": "number"
                },
                "OutputFormat": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                },
                "SerdeInfo": {
                    "$ref": "#/definitions/AWS::Glue::Partition.SerdeInfo"
                },
                "SkewedInfo": {
                    "$ref": "#/definitions/AWS::Glue::Partition.SkewedInfo"
                },
                "SortColumns": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Partition.Order"
                    },
                    "type": "array"
                },
                "StoredAsSubDirectories": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Table": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CatalogId": {
                            "type": "string"
                        },
                        "DatabaseName": {
                            "type": "string"
                        },
                        "TableInput": {
                            "$ref": "#/definitions/AWS::Glue::Table.TableInput"
                        }
                    },
                    "required": [
                        "CatalogId",
                        "DatabaseName",
                        "TableInput"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Table"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Table.Column": {
            "additionalProperties": false,
            "properties": {
                "Comment": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Name"
            ],
            "type": "object"
        },
        "AWS::Glue::Table.Order": {
            "additionalProperties": false,
            "properties": {
                "Column": {
                    "type": "string"
                },
                "SortOrder": {
                    "type": "number"
                }
            },
            "required": [
                "Column",
                "SortOrder"
            ],
            "type": "object"
        },
        "AWS::Glue::Table.SerdeInfo": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                },
                "SerializationLibrary": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Table.SkewedInfo": {
            "additionalProperties": false,
            "properties": {
                "SkewedColumnNames": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SkewedColumnValueLocationMaps": {
                    "type": "object"
                },
                "SkewedColumnValues": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Table.StorageDescriptor": {
            "additionalProperties": false,
            "properties": {
                "BucketColumns": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Columns": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Table.Column"
                    },
                    "type": "array"
                },
                "Compressed": {
                    "type": "boolean"
                },
                "InputFormat": {
                    "type": "string"
                },
                "Location": {
                    "type": "string"
                },
                "NumberOfBuckets": {
                    "type": "number"
                },
                "OutputFormat": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                },
                "SerdeInfo": {
                    "$ref": "#/definitions/AWS::Glue::Table.SerdeInfo"
                },
                "SkewedInfo": {
                    "$ref": "#/definitions/AWS::Glue::Table.SkewedInfo"
                },
                "SortColumns": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Table.Order"
                    },
                    "type": "array"
                },
                "StoredAsSubDirectories": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Table.TableInput": {
            "additionalProperties": false,
            "properties": {
                "Description": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Owner": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                },
                "PartitionKeys": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Table.Column"
                    },
                    "type": "array"
                },
                "Retention": {
                    "type": "number"
                },
                "StorageDescriptor": {
                    "$ref": "#/definitions/AWS::Glue::Table.StorageDescriptor"
                },
                "TableType": {
                    "type": "string"
                },
                "ViewExpandedText": {
                    "type": "string"
                },
                "ViewOriginalText": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Trigger": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Actions": {
                            "items": {
                                "$ref": "#/definitions/AWS::Glue::Trigger.Action"
                            },
                            "type": "array"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Predicate": {
                            "$ref": "#/definitions/AWS::Glue::Trigger.Predicate"
                        },
                        "Schedule": {
                            "type": "string"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Actions",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Glue::Trigger"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Glue::Trigger.Action": {
            "additionalProperties": false,
            "properties": {
                "Arguments": {
                    "type": "object"
                },
                "JobName": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Trigger.Condition": {
            "additionalProperties": false,
            "properties": {
                "JobName": {
                    "type": "string"
                },
                "LogicalOperator": {
                    "type": "string"
                },
                "State": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Glue::Trigger.Predicate": {
            "additionalProperties": false,
            "properties": {
                "Conditions": {
                    "items": {
                        "$ref": "#/definitions/AWS::Glue::Trigger.Condition"
                    },
                    "type": "array"
                },
                "Logical": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::GuardDuty::Detector": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Enable": {
                            "type": "boolean"
                        }
                    },
                    "required": [
                        "Enable"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GuardDuty::Detector"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::GuardDuty::IPSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Activate": {
                            "type": "boolean"
                        },
                        "DetectorId": {
                            "type": "string"
                        },
                        "Format": {
                            "type": "string"
                        },
                        "Location": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Activate",
                        "DetectorId",
                        "Format",
                        "Location"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GuardDuty::IPSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::GuardDuty::Master": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DetectorId": {
                            "type": "string"
                        },
                        "InvitationId": {
                            "type": "string"
                        },
                        "MasterId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DetectorId",
                        "InvitationId",
                        "MasterId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GuardDuty::Master"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::GuardDuty::Member": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DetectorId": {
                            "type": "string"
                        },
                        "Email": {
                            "type": "string"
                        },
                        "MemberId": {
                            "type": "string"
                        },
                        "Message": {
                            "type": "string"
                        },
                        "Status": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DetectorId",
                        "Email",
                        "MemberId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GuardDuty::Member"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::GuardDuty::ThreatIntelSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Activate": {
                            "type": "boolean"
                        },
                        "DetectorId": {
                            "type": "string"
                        },
                        "Format": {
                            "type": "string"
                        },
                        "Location": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Activate",
                        "DetectorId",
                        "Format",
                        "Location"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::GuardDuty::ThreatIntelSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IAM::AccessKey": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Serial": {
                            "type": "number"
                        },
                        "Status": {
                            "type": "string"
                        },
                        "UserName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "UserName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::AccessKey"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IAM::Group": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "GroupName": {
                            "type": "string"
                        },
                        "ManagedPolicyArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Path": {
                            "type": "string"
                        },
                        "Policies": {
                            "items": {
                                "$ref": "#/definitions/AWS::IAM::Group.Policy"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::Group"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::IAM::Group.Policy": {
            "additionalProperties": false,
            "properties": {
                "PolicyDocument": {
                    "type": "object"
                },
                "PolicyName": {
                    "type": "string"
                }
            },
            "required": [
                "PolicyDocument",
                "PolicyName"
            ],
            "type": "object"
        },
        "AWS::IAM::InstanceProfile": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "InstanceProfileName": {
                            "type": "string"
                        },
                        "Path": {
                            "type": "string"
                        },
                        "Roles": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Roles"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::InstanceProfile"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IAM::ManagedPolicy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Groups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ManagedPolicyName": {
                            "type": "string"
                        },
                        "Path": {
                            "type": "string"
                        },
                        "PolicyDocument": {
                            "type": "object"
                        },
                        "Roles": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Users": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "PolicyDocument"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::ManagedPolicy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IAM::Policy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Groups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "PolicyDocument": {
                            "type": "object"
                        },
                        "PolicyName": {
                            "type": "string"
                        },
                        "Roles": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Users": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "PolicyDocument",
                        "PolicyName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::Policy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IAM::Role": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AssumeRolePolicyDocument": {
                            "type": "object"
                        },
                        "ManagedPolicyArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Path": {
                            "type": "string"
                        },
                        "Policies": {
                            "items": {
                                "$ref": "#/definitions/AWS::IAM::Role.Policy"
                            },
                            "type": "array"
                        },
                        "RoleName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AssumeRolePolicyDocument"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::Role"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IAM::Role.Policy": {
            "additionalProperties": false,
            "properties": {
                "PolicyDocument": {
                    "type": "object"
                },
                "PolicyName": {
                    "type": "string"
                }
            },
            "required": [
                "PolicyDocument",
                "PolicyName"
            ],
            "type": "object"
        },
        "AWS::IAM::User": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Groups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "LoginProfile": {
                            "$ref": "#/definitions/AWS::IAM::User.LoginProfile"
                        },
                        "ManagedPolicyArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Path": {
                            "type": "string"
                        },
                        "Policies": {
                            "items": {
                                "$ref": "#/definitions/AWS::IAM::User.Policy"
                            },
                            "type": "array"
                        },
                        "UserName": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::User"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::IAM::User.LoginProfile": {
            "additionalProperties": false,
            "properties": {
                "Password": {
                    "type": "string"
                },
                "PasswordResetRequired": {
                    "type": "boolean"
                }
            },
            "required": [
                "Password"
            ],
            "type": "object"
        },
        "AWS::IAM::User.Policy": {
            "additionalProperties": false,
            "properties": {
                "PolicyDocument": {
                    "type": "object"
                },
                "PolicyName": {
                    "type": "string"
                }
            },
            "required": [
                "PolicyDocument",
                "PolicyName"
            ],
            "type": "object"
        },
        "AWS::IAM::UserToGroupAddition": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "GroupName": {
                            "type": "string"
                        },
                        "Users": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "GroupName",
                        "Users"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IAM::UserToGroupAddition"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Inspector::AssessmentTarget": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AssessmentTargetName": {
                            "type": "string"
                        },
                        "ResourceGroupArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ResourceGroupArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Inspector::AssessmentTarget"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Inspector::AssessmentTemplate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AssessmentTargetArn": {
                            "type": "string"
                        },
                        "AssessmentTemplateName": {
                            "type": "string"
                        },
                        "DurationInSeconds": {
                            "type": "number"
                        },
                        "RulesPackageArns": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "UserAttributesForFindings": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "AssessmentTargetArn",
                        "DurationInSeconds",
                        "RulesPackageArns"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Inspector::AssessmentTemplate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Inspector::ResourceGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ResourceGroupTags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ResourceGroupTags"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Inspector::ResourceGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IoT::Certificate": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CertificateSigningRequest": {
                            "type": "string"
                        },
                        "Status": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "CertificateSigningRequest",
                        "Status"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IoT::Certificate"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IoT::Policy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PolicyDocument": {
                            "type": "object"
                        },
                        "PolicyName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "PolicyDocument"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IoT::Policy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IoT::PolicyPrincipalAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PolicyName": {
                            "type": "string"
                        },
                        "Principal": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "PolicyName",
                        "Principal"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IoT::PolicyPrincipalAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IoT::Thing": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AttributePayload": {
                            "$ref": "#/definitions/AWS::IoT::Thing.AttributePayload"
                        },
                        "ThingName": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IoT::Thing"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::IoT::Thing.AttributePayload": {
            "additionalProperties": false,
            "properties": {
                "Attributes": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "AWS::IoT::ThingPrincipalAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Principal": {
                            "type": "string"
                        },
                        "ThingName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Principal",
                        "ThingName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IoT::ThingPrincipalAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "RuleName": {
                            "type": "string"
                        },
                        "TopicRulePayload": {
                            "$ref": "#/definitions/AWS::IoT::TopicRule.TopicRulePayload"
                        }
                    },
                    "required": [
                        "TopicRulePayload"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::IoT::TopicRule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.Action": {
            "additionalProperties": false,
            "properties": {
                "CloudwatchAlarm": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.CloudwatchAlarmAction"
                },
                "CloudwatchMetric": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.CloudwatchMetricAction"
                },
                "DynamoDB": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.DynamoDBAction"
                },
                "DynamoDBv2": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.DynamoDBv2Action"
                },
                "Elasticsearch": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.ElasticsearchAction"
                },
                "Firehose": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.FirehoseAction"
                },
                "Kinesis": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.KinesisAction"
                },
                "Lambda": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.LambdaAction"
                },
                "Republish": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.RepublishAction"
                },
                "S3": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.S3Action"
                },
                "Sns": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.SnsAction"
                },
                "Sqs": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.SqsAction"
                }
            },
            "type": "object"
        },
        "AWS::IoT::TopicRule.CloudwatchAlarmAction": {
            "additionalProperties": false,
            "properties": {
                "AlarmName": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "StateReason": {
                    "type": "string"
                },
                "StateValue": {
                    "type": "string"
                }
            },
            "required": [
                "AlarmName",
                "RoleArn",
                "StateReason",
                "StateValue"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.CloudwatchMetricAction": {
            "additionalProperties": false,
            "properties": {
                "MetricName": {
                    "type": "string"
                },
                "MetricNamespace": {
                    "type": "string"
                },
                "MetricTimestamp": {
                    "type": "string"
                },
                "MetricUnit": {
                    "type": "string"
                },
                "MetricValue": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                }
            },
            "required": [
                "MetricName",
                "MetricNamespace",
                "MetricUnit",
                "MetricValue",
                "RoleArn"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.DynamoDBAction": {
            "additionalProperties": false,
            "properties": {
                "HashKeyField": {
                    "type": "string"
                },
                "HashKeyType": {
                    "type": "string"
                },
                "HashKeyValue": {
                    "type": "string"
                },
                "PayloadField": {
                    "type": "string"
                },
                "RangeKeyField": {
                    "type": "string"
                },
                "RangeKeyType": {
                    "type": "string"
                },
                "RangeKeyValue": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "TableName": {
                    "type": "string"
                }
            },
            "required": [
                "HashKeyField",
                "HashKeyValue",
                "RoleArn",
                "TableName"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.DynamoDBv2Action": {
            "additionalProperties": false,
            "properties": {
                "PutItem": {
                    "$ref": "#/definitions/AWS::IoT::TopicRule.PutItemInput"
                },
                "RoleArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::IoT::TopicRule.ElasticsearchAction": {
            "additionalProperties": false,
            "properties": {
                "Endpoint": {
                    "type": "string"
                },
                "Id": {
                    "type": "string"
                },
                "Index": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Endpoint",
                "Id",
                "Index",
                "RoleArn",
                "Type"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.FirehoseAction": {
            "additionalProperties": false,
            "properties": {
                "DeliveryStreamName": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "Separator": {
                    "type": "string"
                }
            },
            "required": [
                "DeliveryStreamName",
                "RoleArn"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.KinesisAction": {
            "additionalProperties": false,
            "properties": {
                "PartitionKey": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "StreamName": {
                    "type": "string"
                }
            },
            "required": [
                "RoleArn",
                "StreamName"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.LambdaAction": {
            "additionalProperties": false,
            "properties": {
                "FunctionArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::IoT::TopicRule.PutItemInput": {
            "additionalProperties": false,
            "properties": {
                "TableName": {
                    "type": "string"
                }
            },
            "required": [
                "TableName"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.RepublishAction": {
            "additionalProperties": false,
            "properties": {
                "RoleArn": {
                    "type": "string"
                },
                "Topic": {
                    "type": "string"
                }
            },
            "required": [
                "RoleArn",
                "Topic"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.S3Action": {
            "additionalProperties": false,
            "properties": {
                "BucketName": {
                    "type": "string"
                },
                "Key": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                }
            },
            "required": [
                "BucketName",
                "Key",
                "RoleArn"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.SnsAction": {
            "additionalProperties": false,
            "properties": {
                "MessageFormat": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "TargetArn": {
                    "type": "string"
                }
            },
            "required": [
                "RoleArn",
                "TargetArn"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.SqsAction": {
            "additionalProperties": false,
            "properties": {
                "QueueUrl": {
                    "type": "string"
                },
                "RoleArn": {
                    "type": "string"
                },
                "UseBase64": {
                    "type": "boolean"
                }
            },
            "required": [
                "QueueUrl",
                "RoleArn"
            ],
            "type": "object"
        },
        "AWS::IoT::TopicRule.TopicRulePayload": {
            "additionalProperties": false,
            "properties": {
                "Actions": {
                    "items": {
                        "$ref": "#/definitions/AWS::IoT::TopicRule.Action"
                    },
                    "type": "array"
                },
                "AwsIotSqlVersion": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "RuleDisabled": {
                    "type": "boolean"
                },
                "Sql": {
                    "type": "string"
                }
            },
            "required": [
                "Actions",
                "RuleDisabled",
                "Sql"
            ],
            "type": "object"
        },
        "AWS::KMS::Alias": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AliasName": {
                            "type": "string"
                        },
                        "TargetKeyId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "AliasName",
                        "TargetKeyId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::KMS::Alias"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::KMS::Key": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "EnableKeyRotation": {
                            "type": "boolean"
                        },
                        "Enabled": {
                            "type": "boolean"
                        },
                        "KeyPolicy": {
                            "type": "object"
                        },
                        "KeyUsage": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "KeyPolicy"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::KMS::Key"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Kinesis::Stream": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "RetentionPeriodHours": {
                            "type": "number"
                        },
                        "ShardCount": {
                            "type": "number"
                        },
                        "StreamEncryption": {
                            "$ref": "#/definitions/AWS::Kinesis::Stream.StreamEncryption"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ShardCount"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Kinesis::Stream"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Kinesis::Stream.StreamEncryption": {
            "additionalProperties": false,
            "properties": {
                "EncryptionType": {
                    "type": "string"
                },
                "KeyId": {
                    "type": "string"
                }
            },
            "required": [
                "EncryptionType",
                "KeyId"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationCode": {
                            "type": "string"
                        },
                        "ApplicationDescription": {
                            "type": "string"
                        },
                        "ApplicationName": {
                            "type": "string"
                        },
                        "Inputs": {
                            "items": {
                                "$ref": "#/definitions/AWS::KinesisAnalytics::Application.Input"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Inputs"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::KinesisAnalytics::Application"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.CSVMappingParameters": {
            "additionalProperties": false,
            "properties": {
                "RecordColumnDelimiter": {
                    "type": "string"
                },
                "RecordRowDelimiter": {
                    "type": "string"
                }
            },
            "required": [
                "RecordColumnDelimiter",
                "RecordRowDelimiter"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.Input": {
            "additionalProperties": false,
            "properties": {
                "InputParallelism": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.InputParallelism"
                },
                "InputProcessingConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.InputProcessingConfiguration"
                },
                "InputSchema": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.InputSchema"
                },
                "KinesisFirehoseInput": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.KinesisFirehoseInput"
                },
                "KinesisStreamsInput": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.KinesisStreamsInput"
                },
                "NamePrefix": {
                    "type": "string"
                }
            },
            "required": [
                "InputSchema",
                "NamePrefix"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.InputLambdaProcessor": {
            "additionalProperties": false,
            "properties": {
                "ResourceARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "ResourceARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.InputParallelism": {
            "additionalProperties": false,
            "properties": {
                "Count": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.InputProcessingConfiguration": {
            "additionalProperties": false,
            "properties": {
                "InputLambdaProcessor": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.InputLambdaProcessor"
                }
            },
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.InputSchema": {
            "additionalProperties": false,
            "properties": {
                "RecordColumns": {
                    "items": {
                        "$ref": "#/definitions/AWS::KinesisAnalytics::Application.RecordColumn"
                    },
                    "type": "array"
                },
                "RecordEncoding": {
                    "type": "string"
                },
                "RecordFormat": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.RecordFormat"
                }
            },
            "required": [
                "RecordColumns",
                "RecordFormat"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.JSONMappingParameters": {
            "additionalProperties": false,
            "properties": {
                "RecordRowPath": {
                    "type": "string"
                }
            },
            "required": [
                "RecordRowPath"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.KinesisFirehoseInput": {
            "additionalProperties": false,
            "properties": {
                "ResourceARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "ResourceARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.KinesisStreamsInput": {
            "additionalProperties": false,
            "properties": {
                "ResourceARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "ResourceARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.MappingParameters": {
            "additionalProperties": false,
            "properties": {
                "CSVMappingParameters": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.CSVMappingParameters"
                },
                "JSONMappingParameters": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.JSONMappingParameters"
                }
            },
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.RecordColumn": {
            "additionalProperties": false,
            "properties": {
                "Mapping": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "SqlType": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "SqlType"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::Application.RecordFormat": {
            "additionalProperties": false,
            "properties": {
                "MappingParameters": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::Application.MappingParameters"
                },
                "RecordFormatType": {
                    "type": "string"
                }
            },
            "required": [
                "RecordFormatType"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationOutput": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "Output": {
                            "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationOutput.Output"
                        }
                    },
                    "required": [
                        "ApplicationName",
                        "Output"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::KinesisAnalytics::ApplicationOutput"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema": {
            "additionalProperties": false,
            "properties": {
                "RecordFormatType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationOutput.KinesisFirehoseOutput": {
            "additionalProperties": false,
            "properties": {
                "ResourceARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "ResourceARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationOutput.KinesisStreamsOutput": {
            "additionalProperties": false,
            "properties": {
                "ResourceARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "ResourceARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationOutput.LambdaOutput": {
            "additionalProperties": false,
            "properties": {
                "ResourceARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "ResourceARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationOutput.Output": {
            "additionalProperties": false,
            "properties": {
                "DestinationSchema": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema"
                },
                "KinesisFirehoseOutput": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationOutput.KinesisFirehoseOutput"
                },
                "KinesisStreamsOutput": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationOutput.KinesisStreamsOutput"
                },
                "LambdaOutput": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationOutput.LambdaOutput"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "DestinationSchema"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApplicationName": {
                            "type": "string"
                        },
                        "ReferenceDataSource": {
                            "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource"
                        }
                    },
                    "required": [
                        "ApplicationName",
                        "ReferenceDataSource"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::KinesisAnalytics::ApplicationReferenceDataSource"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.CSVMappingParameters": {
            "additionalProperties": false,
            "properties": {
                "RecordColumnDelimiter": {
                    "type": "string"
                },
                "RecordRowDelimiter": {
                    "type": "string"
                }
            },
            "required": [
                "RecordColumnDelimiter",
                "RecordRowDelimiter"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.JSONMappingParameters": {
            "additionalProperties": false,
            "properties": {
                "RecordRowPath": {
                    "type": "string"
                }
            },
            "required": [
                "RecordRowPath"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.MappingParameters": {
            "additionalProperties": false,
            "properties": {
                "CSVMappingParameters": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.CSVMappingParameters"
                },
                "JSONMappingParameters": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.JSONMappingParameters"
                }
            },
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordColumn": {
            "additionalProperties": false,
            "properties": {
                "Mapping": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "SqlType": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "SqlType"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordFormat": {
            "additionalProperties": false,
            "properties": {
                "MappingParameters": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.MappingParameters"
                },
                "RecordFormatType": {
                    "type": "string"
                }
            },
            "required": [
                "RecordFormatType"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource": {
            "additionalProperties": false,
            "properties": {
                "ReferenceSchema": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceSchema"
                },
                "S3ReferenceDataSource": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.S3ReferenceDataSource"
                },
                "TableName": {
                    "type": "string"
                }
            },
            "required": [
                "ReferenceSchema"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceSchema": {
            "additionalProperties": false,
            "properties": {
                "RecordColumns": {
                    "items": {
                        "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordColumn"
                    },
                    "type": "array"
                },
                "RecordEncoding": {
                    "type": "string"
                },
                "RecordFormat": {
                    "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordFormat"
                }
            },
            "required": [
                "RecordColumns",
                "RecordFormat"
            ],
            "type": "object"
        },
        "AWS::KinesisAnalytics::ApplicationReferenceDataSource.S3ReferenceDataSource": {
            "additionalProperties": false,
            "properties": {
                "BucketARN": {
                    "type": "string"
                },
                "FileKey": {
                    "type": "string"
                },
                "ReferenceRoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "BucketARN",
                "FileKey",
                "ReferenceRoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DeliveryStreamName": {
                            "type": "string"
                        },
                        "DeliveryStreamType": {
                            "type": "string"
                        },
                        "ElasticsearchDestinationConfiguration": {
                            "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ElasticsearchDestinationConfiguration"
                        },
                        "ExtendedS3DestinationConfiguration": {
                            "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ExtendedS3DestinationConfiguration"
                        },
                        "KinesisStreamSourceConfiguration": {
                            "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.KinesisStreamSourceConfiguration"
                        },
                        "RedshiftDestinationConfiguration": {
                            "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.RedshiftDestinationConfiguration"
                        },
                        "S3DestinationConfiguration": {
                            "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::KinesisFirehose::DeliveryStream"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.BufferingHints": {
            "additionalProperties": false,
            "properties": {
                "IntervalInSeconds": {
                    "type": "number"
                },
                "SizeInMBs": {
                    "type": "number"
                }
            },
            "required": [
                "IntervalInSeconds",
                "SizeInMBs"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions": {
            "additionalProperties": false,
            "properties": {
                "Enabled": {
                    "type": "boolean"
                },
                "LogGroupName": {
                    "type": "string"
                },
                "LogStreamName": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.CopyCommand": {
            "additionalProperties": false,
            "properties": {
                "CopyOptions": {
                    "type": "string"
                },
                "DataTableColumns": {
                    "type": "string"
                },
                "DataTableName": {
                    "type": "string"
                }
            },
            "required": [
                "DataTableName"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints": {
            "additionalProperties": false,
            "properties": {
                "IntervalInSeconds": {
                    "type": "number"
                },
                "SizeInMBs": {
                    "type": "number"
                }
            },
            "required": [
                "IntervalInSeconds",
                "SizeInMBs"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.ElasticsearchDestinationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "BufferingHints": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints"
                },
                "CloudWatchLoggingOptions": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions"
                },
                "DomainARN": {
                    "type": "string"
                },
                "IndexName": {
                    "type": "string"
                },
                "IndexRotationPeriod": {
                    "type": "string"
                },
                "ProcessingConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ProcessingConfiguration"
                },
                "RetryOptions": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions"
                },
                "RoleARN": {
                    "type": "string"
                },
                "S3BackupMode": {
                    "type": "string"
                },
                "S3Configuration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration"
                },
                "TypeName": {
                    "type": "string"
                }
            },
            "required": [
                "BufferingHints",
                "DomainARN",
                "IndexName",
                "IndexRotationPeriod",
                "RetryOptions",
                "RoleARN",
                "S3BackupMode",
                "S3Configuration",
                "TypeName"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions": {
            "additionalProperties": false,
            "properties": {
                "DurationInSeconds": {
                    "type": "number"
                }
            },
            "required": [
                "DurationInSeconds"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration": {
            "additionalProperties": false,
            "properties": {
                "KMSEncryptionConfig": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig"
                },
                "NoEncryptionConfig": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.ExtendedS3DestinationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "BucketARN": {
                    "type": "string"
                },
                "BufferingHints": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.BufferingHints"
                },
                "CloudWatchLoggingOptions": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions"
                },
                "CompressionFormat": {
                    "type": "string"
                },
                "EncryptionConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration"
                },
                "Prefix": {
                    "type": "string"
                },
                "ProcessingConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ProcessingConfiguration"
                },
                "RoleARN": {
                    "type": "string"
                },
                "S3BackupConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration"
                },
                "S3BackupMode": {
                    "type": "string"
                }
            },
            "required": [
                "BucketARN",
                "BufferingHints",
                "CompressionFormat",
                "Prefix",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig": {
            "additionalProperties": false,
            "properties": {
                "AWSKMSKeyARN": {
                    "type": "string"
                }
            },
            "required": [
                "AWSKMSKeyARN"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.KinesisStreamSourceConfiguration": {
            "additionalProperties": false,
            "properties": {
                "KinesisStreamARN": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "KinesisStreamARN",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.ProcessingConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Enabled": {
                    "type": "boolean"
                },
                "Processors": {
                    "items": {
                        "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.Processor"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Enabled",
                "Processors"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.Processor": {
            "additionalProperties": false,
            "properties": {
                "Parameters": {
                    "items": {
                        "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ProcessorParameter"
                    },
                    "type": "array"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Parameters",
                "Type"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.ProcessorParameter": {
            "additionalProperties": false,
            "properties": {
                "ParameterName": {
                    "type": "string"
                },
                "ParameterValue": {
                    "type": "string"
                }
            },
            "required": [
                "ParameterName",
                "ParameterValue"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.RedshiftDestinationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "CloudWatchLoggingOptions": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions"
                },
                "ClusterJDBCURL": {
                    "type": "string"
                },
                "CopyCommand": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.CopyCommand"
                },
                "Password": {
                    "type": "string"
                },
                "ProcessingConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.ProcessingConfiguration"
                },
                "RoleARN": {
                    "type": "string"
                },
                "S3Configuration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration"
                },
                "Username": {
                    "type": "string"
                }
            },
            "required": [
                "ClusterJDBCURL",
                "CopyCommand",
                "Password",
                "RoleARN",
                "S3Configuration",
                "Username"
            ],
            "type": "object"
        },
        "AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "BucketARN": {
                    "type": "string"
                },
                "BufferingHints": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.BufferingHints"
                },
                "CloudWatchLoggingOptions": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions"
                },
                "CompressionFormat": {
                    "type": "string"
                },
                "EncryptionConfiguration": {
                    "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration"
                },
                "Prefix": {
                    "type": "string"
                },
                "RoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "BucketARN",
                "BufferingHints",
                "CompressionFormat",
                "RoleARN"
            ],
            "type": "object"
        },
        "AWS::Lambda::Alias": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "FunctionName": {
                            "type": "string"
                        },
                        "FunctionVersion": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "RoutingConfig": {
                            "$ref": "#/definitions/AWS::Lambda::Alias.AliasRoutingConfiguration"
                        }
                    },
                    "required": [
                        "FunctionName",
                        "FunctionVersion",
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Lambda::Alias"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Lambda::Alias.AliasRoutingConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AdditionalVersionWeights": {
                    "items": {
                        "$ref": "#/definitions/AWS::Lambda::Alias.VersionWeight"
                    },
                    "type": "array"
                }
            },
            "required": [
                "AdditionalVersionWeights"
            ],
            "type": "object"
        },
        "AWS::Lambda::Alias.VersionWeight": {
            "additionalProperties": false,
            "properties": {
                "FunctionVersion": {
                    "type": "string"
                },
                "FunctionWeight": {
                    "type": "number"
                }
            },
            "required": [
                "FunctionVersion",
                "FunctionWeight"
            ],
            "type": "object"
        },
        "AWS::Lambda::EventSourceMapping": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "BatchSize": {
                            "type": "number"
                        },
                        "Enabled": {
                            "type": "boolean"
                        },
                        "EventSourceArn": {
                            "type": "string"
                        },
                        "FunctionName": {
                            "type": "string"
                        },
                        "StartingPosition": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "EventSourceArn",
                        "FunctionName",
                        "StartingPosition"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Lambda::EventSourceMapping"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Lambda::Function": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Code": {
                            "$ref": "#/definitions/AWS::Lambda::Function.Code"
                        },
                        "DeadLetterConfig": {
                            "$ref": "#/definitions/AWS::Lambda::Function.DeadLetterConfig"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Environment": {
                            "$ref": "#/definitions/AWS::Lambda::Function.Environment"
                        },
                        "FunctionName": {
                            "type": "string"
                        },
                        "Handler": {
                            "type": "string"
                        },
                        "KmsKeyArn": {
                            "type": "string"
                        },
                        "MemorySize": {
                            "type": "number"
                        },
                        "ReservedConcurrentExecutions": {
                            "type": "number"
                        },
                        "Role": {
                            "type": "string"
                        },
                        "Runtime": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Timeout": {
                            "type": "number"
                        },
                        "TracingConfig": {
                            "$ref": "#/definitions/AWS::Lambda::Function.TracingConfig"
                        },
                        "VpcConfig": {
                            "$ref": "#/definitions/AWS::Lambda::Function.VpcConfig"
                        }
                    },
                    "required": [
                        "Code",
                        "Handler",
                        "Role",
                        "Runtime"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Lambda::Function"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Lambda::Function.Code": {
            "additionalProperties": false,
            "properties": {
                "S3Bucket": {
                    "type": "string"
                },
                "S3Key": {
                    "type": "string"
                },
                "S3ObjectVersion": {
                    "type": "string"
                },
                "ZipFile": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Lambda::Function.DeadLetterConfig": {
            "additionalProperties": false,
            "properties": {
                "TargetArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Lambda::Function.Environment": {
            "additionalProperties": false,
            "properties": {
                "Variables": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "AWS::Lambda::Function.TracingConfig": {
            "additionalProperties": false,
            "properties": {
                "Mode": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Lambda::Function.VpcConfig": {
            "additionalProperties": false,
            "properties": {
                "SecurityGroupIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SubnetIds": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "SecurityGroupIds",
                "SubnetIds"
            ],
            "type": "object"
        },
        "AWS::Lambda::Permission": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Action": {
                            "type": "string"
                        },
                        "EventSourceToken": {
                            "type": "string"
                        },
                        "FunctionName": {
                            "type": "string"
                        },
                        "Principal": {
                            "type": "string"
                        },
                        "SourceAccount": {
                            "type": "string"
                        },
                        "SourceArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Action",
                        "FunctionName",
                        "Principal"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Lambda::Permission"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Lambda::Version": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CodeSha256": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "FunctionName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "FunctionName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Lambda::Version"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Logs::Destination": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DestinationName": {
                            "type": "string"
                        },
                        "DestinationPolicy": {
                            "type": "string"
                        },
                        "RoleArn": {
                            "type": "string"
                        },
                        "TargetArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DestinationName",
                        "DestinationPolicy",
                        "RoleArn",
                        "TargetArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Logs::Destination"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Logs::LogGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "LogGroupName": {
                            "type": "string"
                        },
                        "RetentionInDays": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Logs::LogGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Logs::LogStream": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "LogGroupName": {
                            "type": "string"
                        },
                        "LogStreamName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "LogGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Logs::LogStream"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Logs::MetricFilter": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "FilterPattern": {
                            "type": "string"
                        },
                        "LogGroupName": {
                            "type": "string"
                        },
                        "MetricTransformations": {
                            "items": {
                                "$ref": "#/definitions/AWS::Logs::MetricFilter.MetricTransformation"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "FilterPattern",
                        "LogGroupName",
                        "MetricTransformations"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Logs::MetricFilter"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Logs::MetricFilter.MetricTransformation": {
            "additionalProperties": false,
            "properties": {
                "MetricName": {
                    "type": "string"
                },
                "MetricNamespace": {
                    "type": "string"
                },
                "MetricValue": {
                    "type": "string"
                }
            },
            "required": [
                "MetricName",
                "MetricNamespace",
                "MetricValue"
            ],
            "type": "object"
        },
        "AWS::Logs::SubscriptionFilter": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DestinationArn": {
                            "type": "string"
                        },
                        "FilterPattern": {
                            "type": "string"
                        },
                        "LogGroupName": {
                            "type": "string"
                        },
                        "RoleArn": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DestinationArn",
                        "FilterPattern",
                        "LogGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Logs::SubscriptionFilter"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::App": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AppSource": {
                            "$ref": "#/definitions/AWS::OpsWorks::App.Source"
                        },
                        "Attributes": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "DataSources": {
                            "items": {
                                "$ref": "#/definitions/AWS::OpsWorks::App.DataSource"
                            },
                            "type": "array"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Domains": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "EnableSsl": {
                            "type": "boolean"
                        },
                        "Environment": {
                            "items": {
                                "$ref": "#/definitions/AWS::OpsWorks::App.EnvironmentVariable"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Shortname": {
                            "type": "string"
                        },
                        "SslConfiguration": {
                            "$ref": "#/definitions/AWS::OpsWorks::App.SslConfiguration"
                        },
                        "StackId": {
                            "type": "string"
                        },
                        "Type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name",
                        "StackId",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::App"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::App.DataSource": {
            "additionalProperties": false,
            "properties": {
                "Arn": {
                    "type": "string"
                },
                "DatabaseName": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::App.EnvironmentVariable": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Secure": {
                    "type": "boolean"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::App.Source": {
            "additionalProperties": false,
            "properties": {
                "Password": {
                    "type": "string"
                },
                "Revision": {
                    "type": "string"
                },
                "SshKey": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Url": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::App.SslConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Certificate": {
                    "type": "string"
                },
                "Chain": {
                    "type": "string"
                },
                "PrivateKey": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::ElasticLoadBalancerAttachment": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ElasticLoadBalancerName": {
                            "type": "string"
                        },
                        "LayerId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ElasticLoadBalancerName",
                        "LayerId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::ElasticLoadBalancerAttachment"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Instance": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AgentVersion": {
                            "type": "string"
                        },
                        "AmiId": {
                            "type": "string"
                        },
                        "Architecture": {
                            "type": "string"
                        },
                        "AutoScalingType": {
                            "type": "string"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "BlockDeviceMappings": {
                            "items": {
                                "$ref": "#/definitions/AWS::OpsWorks::Instance.BlockDeviceMapping"
                            },
                            "type": "array"
                        },
                        "EbsOptimized": {
                            "type": "boolean"
                        },
                        "ElasticIps": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Hostname": {
                            "type": "string"
                        },
                        "InstallUpdatesOnBoot": {
                            "type": "boolean"
                        },
                        "InstanceType": {
                            "type": "string"
                        },
                        "LayerIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Os": {
                            "type": "string"
                        },
                        "RootDeviceType": {
                            "type": "string"
                        },
                        "SshKeyName": {
                            "type": "string"
                        },
                        "StackId": {
                            "type": "string"
                        },
                        "SubnetId": {
                            "type": "string"
                        },
                        "Tenancy": {
                            "type": "string"
                        },
                        "TimeBasedAutoScaling": {
                            "$ref": "#/definitions/AWS::OpsWorks::Instance.TimeBasedAutoScaling"
                        },
                        "VirtualizationType": {
                            "type": "string"
                        },
                        "Volumes": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "InstanceType",
                        "LayerIds",
                        "StackId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::Instance"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Instance.BlockDeviceMapping": {
            "additionalProperties": false,
            "properties": {
                "DeviceName": {
                    "type": "string"
                },
                "Ebs": {
                    "$ref": "#/definitions/AWS::OpsWorks::Instance.EbsBlockDevice"
                },
                "NoDevice": {
                    "type": "string"
                },
                "VirtualName": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Instance.EbsBlockDevice": {
            "additionalProperties": false,
            "properties": {
                "DeleteOnTermination": {
                    "type": "boolean"
                },
                "Iops": {
                    "type": "number"
                },
                "SnapshotId": {
                    "type": "string"
                },
                "VolumeSize": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Instance.TimeBasedAutoScaling": {
            "additionalProperties": false,
            "properties": {
                "Friday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Monday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Saturday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Sunday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Thursday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Tuesday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Wednesday": {
                    "additionalProperties": false,
                    "patternProperties": {
                        "^[a-zA-Z0-9]+$": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Layer": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Attributes": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "AutoAssignElasticIps": {
                            "type": "boolean"
                        },
                        "AutoAssignPublicIps": {
                            "type": "boolean"
                        },
                        "CustomInstanceProfileArn": {
                            "type": "string"
                        },
                        "CustomJson": {
                            "type": "object"
                        },
                        "CustomRecipes": {
                            "$ref": "#/definitions/AWS::OpsWorks::Layer.Recipes"
                        },
                        "CustomSecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "EnableAutoHealing": {
                            "type": "boolean"
                        },
                        "InstallUpdatesOnBoot": {
                            "type": "boolean"
                        },
                        "LifecycleEventConfiguration": {
                            "$ref": "#/definitions/AWS::OpsWorks::Layer.LifecycleEventConfiguration"
                        },
                        "LoadBasedAutoScaling": {
                            "$ref": "#/definitions/AWS::OpsWorks::Layer.LoadBasedAutoScaling"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Packages": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Shortname": {
                            "type": "string"
                        },
                        "StackId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Type": {
                            "type": "string"
                        },
                        "UseEbsOptimizedInstances": {
                            "type": "boolean"
                        },
                        "VolumeConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::OpsWorks::Layer.VolumeConfiguration"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "AutoAssignElasticIps",
                        "AutoAssignPublicIps",
                        "EnableAutoHealing",
                        "Name",
                        "Shortname",
                        "StackId",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::Layer"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Layer.AutoScalingThresholds": {
            "additionalProperties": false,
            "properties": {
                "CpuThreshold": {
                    "type": "number"
                },
                "IgnoreMetricsTime": {
                    "type": "number"
                },
                "InstanceCount": {
                    "type": "number"
                },
                "LoadThreshold": {
                    "type": "number"
                },
                "MemoryThreshold": {
                    "type": "number"
                },
                "ThresholdsWaitTime": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Layer.LifecycleEventConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ShutdownEventConfiguration": {
                    "$ref": "#/definitions/AWS::OpsWorks::Layer.ShutdownEventConfiguration"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Layer.LoadBasedAutoScaling": {
            "additionalProperties": false,
            "properties": {
                "DownScaling": {
                    "$ref": "#/definitions/AWS::OpsWorks::Layer.AutoScalingThresholds"
                },
                "Enable": {
                    "type": "boolean"
                },
                "UpScaling": {
                    "$ref": "#/definitions/AWS::OpsWorks::Layer.AutoScalingThresholds"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Layer.Recipes": {
            "additionalProperties": false,
            "properties": {
                "Configure": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Deploy": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Setup": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Shutdown": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Undeploy": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Layer.ShutdownEventConfiguration": {
            "additionalProperties": false,
            "properties": {
                "DelayUntilElbConnectionsDrained": {
                    "type": "boolean"
                },
                "ExecutionTimeout": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Layer.VolumeConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Iops": {
                    "type": "number"
                },
                "MountPoint": {
                    "type": "string"
                },
                "NumberOfDisks": {
                    "type": "number"
                },
                "RaidLevel": {
                    "type": "number"
                },
                "Size": {
                    "type": "number"
                },
                "VolumeType": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Stack": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AgentVersion": {
                            "type": "string"
                        },
                        "Attributes": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "ChefConfiguration": {
                            "$ref": "#/definitions/AWS::OpsWorks::Stack.ChefConfiguration"
                        },
                        "CloneAppIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ClonePermissions": {
                            "type": "boolean"
                        },
                        "ConfigurationManager": {
                            "$ref": "#/definitions/AWS::OpsWorks::Stack.StackConfigurationManager"
                        },
                        "CustomCookbooksSource": {
                            "$ref": "#/definitions/AWS::OpsWorks::Stack.Source"
                        },
                        "CustomJson": {
                            "type": "object"
                        },
                        "DefaultAvailabilityZone": {
                            "type": "string"
                        },
                        "DefaultInstanceProfileArn": {
                            "type": "string"
                        },
                        "DefaultOs": {
                            "type": "string"
                        },
                        "DefaultRootDeviceType": {
                            "type": "string"
                        },
                        "DefaultSshKeyName": {
                            "type": "string"
                        },
                        "DefaultSubnetId": {
                            "type": "string"
                        },
                        "EcsClusterArn": {
                            "type": "string"
                        },
                        "ElasticIps": {
                            "items": {
                                "$ref": "#/definitions/AWS::OpsWorks::Stack.ElasticIp"
                            },
                            "type": "array"
                        },
                        "HostnameTheme": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "RdsDbInstances": {
                            "items": {
                                "$ref": "#/definitions/AWS::OpsWorks::Stack.RdsDbInstance"
                            },
                            "type": "array"
                        },
                        "ServiceRoleArn": {
                            "type": "string"
                        },
                        "SourceStackId": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "UseCustomCookbooks": {
                            "type": "boolean"
                        },
                        "UseOpsworksSecurityGroups": {
                            "type": "boolean"
                        },
                        "VpcId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DefaultInstanceProfileArn",
                        "Name",
                        "ServiceRoleArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::Stack"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Stack.ChefConfiguration": {
            "additionalProperties": false,
            "properties": {
                "BerkshelfVersion": {
                    "type": "string"
                },
                "ManageBerkshelf": {
                    "type": "boolean"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Stack.ElasticIp": {
            "additionalProperties": false,
            "properties": {
                "Ip": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "Ip"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Stack.RdsDbInstance": {
            "additionalProperties": false,
            "properties": {
                "DbPassword": {
                    "type": "string"
                },
                "DbUser": {
                    "type": "string"
                },
                "RdsDbInstanceArn": {
                    "type": "string"
                }
            },
            "required": [
                "DbPassword",
                "DbUser",
                "RdsDbInstanceArn"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Stack.Source": {
            "additionalProperties": false,
            "properties": {
                "Password": {
                    "type": "string"
                },
                "Revision": {
                    "type": "string"
                },
                "SshKey": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Url": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::Stack.StackConfigurationManager": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::OpsWorks::UserProfile": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllowSelfManagement": {
                            "type": "boolean"
                        },
                        "IamUserArn": {
                            "type": "string"
                        },
                        "SshPublicKey": {
                            "type": "string"
                        },
                        "SshUsername": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "IamUserArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::UserProfile"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::OpsWorks::Volume": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Ec2VolumeId": {
                            "type": "string"
                        },
                        "MountPoint": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "StackId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Ec2VolumeId",
                        "StackId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::OpsWorks::Volume"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBCluster": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AvailabilityZones": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "BackupRetentionPeriod": {
                            "type": "number"
                        },
                        "DBClusterIdentifier": {
                            "type": "string"
                        },
                        "DBClusterParameterGroupName": {
                            "type": "string"
                        },
                        "DBSubnetGroupName": {
                            "type": "string"
                        },
                        "DatabaseName": {
                            "type": "string"
                        },
                        "Engine": {
                            "type": "string"
                        },
                        "EngineVersion": {
                            "type": "string"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "MasterUserPassword": {
                            "type": "string"
                        },
                        "MasterUsername": {
                            "type": "string"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "PreferredBackupWindow": {
                            "type": "string"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "ReplicationSourceIdentifier": {
                            "type": "string"
                        },
                        "SnapshotIdentifier": {
                            "type": "string"
                        },
                        "StorageEncrypted": {
                            "type": "boolean"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcSecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Engine"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBCluster"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBClusterParameterGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Family": {
                            "type": "string"
                        },
                        "Parameters": {
                            "type": "object"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Description",
                        "Family",
                        "Parameters"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBClusterParameterGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBInstance": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllocatedStorage": {
                            "type": "string"
                        },
                        "AllowMajorVersionUpgrade": {
                            "type": "boolean"
                        },
                        "AutoMinorVersionUpgrade": {
                            "type": "boolean"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "BackupRetentionPeriod": {
                            "type": "string"
                        },
                        "CharacterSetName": {
                            "type": "string"
                        },
                        "CopyTagsToSnapshot": {
                            "type": "boolean"
                        },
                        "DBClusterIdentifier": {
                            "type": "string"
                        },
                        "DBInstanceClass": {
                            "type": "string"
                        },
                        "DBInstanceIdentifier": {
                            "type": "string"
                        },
                        "DBName": {
                            "type": "string"
                        },
                        "DBParameterGroupName": {
                            "type": "string"
                        },
                        "DBSecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "DBSnapshotIdentifier": {
                            "type": "string"
                        },
                        "DBSubnetGroupName": {
                            "type": "string"
                        },
                        "Domain": {
                            "type": "string"
                        },
                        "DomainIAMRoleName": {
                            "type": "string"
                        },
                        "Engine": {
                            "type": "string"
                        },
                        "EngineVersion": {
                            "type": "string"
                        },
                        "Iops": {
                            "type": "number"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "LicenseModel": {
                            "type": "string"
                        },
                        "MasterUserPassword": {
                            "type": "string"
                        },
                        "MasterUsername": {
                            "type": "string"
                        },
                        "MonitoringInterval": {
                            "type": "number"
                        },
                        "MonitoringRoleArn": {
                            "type": "string"
                        },
                        "MultiAZ": {
                            "type": "boolean"
                        },
                        "OptionGroupName": {
                            "type": "string"
                        },
                        "Port": {
                            "type": "string"
                        },
                        "PreferredBackupWindow": {
                            "type": "string"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "PubliclyAccessible": {
                            "type": "boolean"
                        },
                        "SourceDBInstanceIdentifier": {
                            "type": "string"
                        },
                        "SourceRegion": {
                            "type": "string"
                        },
                        "StorageEncrypted": {
                            "type": "boolean"
                        },
                        "StorageType": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "Timezone": {
                            "type": "string"
                        },
                        "VPCSecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DBInstanceClass"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBInstance"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBParameterGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Family": {
                            "type": "string"
                        },
                        "Parameters": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Description",
                        "Family"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBParameterGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBSecurityGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DBSecurityGroupIngress": {
                            "items": {
                                "$ref": "#/definitions/AWS::RDS::DBSecurityGroup.Ingress"
                            },
                            "type": "array"
                        },
                        "EC2VpcId": {
                            "type": "string"
                        },
                        "GroupDescription": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DBSecurityGroupIngress",
                        "GroupDescription"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBSecurityGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBSecurityGroup.Ingress": {
            "additionalProperties": false,
            "properties": {
                "CIDRIP": {
                    "type": "string"
                },
                "EC2SecurityGroupId": {
                    "type": "string"
                },
                "EC2SecurityGroupName": {
                    "type": "string"
                },
                "EC2SecurityGroupOwnerId": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::RDS::DBSecurityGroupIngress": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CIDRIP": {
                            "type": "string"
                        },
                        "DBSecurityGroupName": {
                            "type": "string"
                        },
                        "EC2SecurityGroupId": {
                            "type": "string"
                        },
                        "EC2SecurityGroupName": {
                            "type": "string"
                        },
                        "EC2SecurityGroupOwnerId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DBSecurityGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBSecurityGroupIngress"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::DBSubnetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DBSubnetGroupDescription": {
                            "type": "string"
                        },
                        "DBSubnetGroupName": {
                            "type": "string"
                        },
                        "SubnetIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DBSubnetGroupDescription",
                        "SubnetIds"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::DBSubnetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::EventSubscription": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Enabled": {
                            "type": "boolean"
                        },
                        "EventCategories": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SnsTopicArn": {
                            "type": "string"
                        },
                        "SourceIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SourceType": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "SnsTopicArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::EventSubscription"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::OptionGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "EngineName": {
                            "type": "string"
                        },
                        "MajorEngineVersion": {
                            "type": "string"
                        },
                        "OptionConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::RDS::OptionGroup.OptionConfiguration"
                            },
                            "type": "array"
                        },
                        "OptionGroupDescription": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "EngineName",
                        "MajorEngineVersion",
                        "OptionConfigurations",
                        "OptionGroupDescription"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::RDS::OptionGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::RDS::OptionGroup.OptionConfiguration": {
            "additionalProperties": false,
            "properties": {
                "DBSecurityGroupMemberships": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "OptionName": {
                    "type": "string"
                },
                "OptionSettings": {
                    "$ref": "#/definitions/AWS::RDS::OptionGroup.OptionSetting"
                },
                "OptionVersion": {
                    "type": "string"
                },
                "Port": {
                    "type": "number"
                },
                "VpcSecurityGroupMemberships": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "OptionName"
            ],
            "type": "object"
        },
        "AWS::RDS::OptionGroup.OptionSetting": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Redshift::Cluster": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllowVersionUpgrade": {
                            "type": "boolean"
                        },
                        "AutomatedSnapshotRetentionPeriod": {
                            "type": "number"
                        },
                        "AvailabilityZone": {
                            "type": "string"
                        },
                        "ClusterIdentifier": {
                            "type": "string"
                        },
                        "ClusterParameterGroupName": {
                            "type": "string"
                        },
                        "ClusterSecurityGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ClusterSubnetGroupName": {
                            "type": "string"
                        },
                        "ClusterType": {
                            "type": "string"
                        },
                        "ClusterVersion": {
                            "type": "string"
                        },
                        "DBName": {
                            "type": "string"
                        },
                        "ElasticIp": {
                            "type": "string"
                        },
                        "Encrypted": {
                            "type": "boolean"
                        },
                        "HsmClientCertificateIdentifier": {
                            "type": "string"
                        },
                        "HsmConfigurationIdentifier": {
                            "type": "string"
                        },
                        "IamRoles": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "KmsKeyId": {
                            "type": "string"
                        },
                        "LoggingProperties": {
                            "$ref": "#/definitions/AWS::Redshift::Cluster.LoggingProperties"
                        },
                        "MasterUserPassword": {
                            "type": "string"
                        },
                        "MasterUsername": {
                            "type": "string"
                        },
                        "NodeType": {
                            "type": "string"
                        },
                        "NumberOfNodes": {
                            "type": "number"
                        },
                        "OwnerAccount": {
                            "type": "string"
                        },
                        "Port": {
                            "type": "number"
                        },
                        "PreferredMaintenanceWindow": {
                            "type": "string"
                        },
                        "PubliclyAccessible": {
                            "type": "boolean"
                        },
                        "SnapshotClusterIdentifier": {
                            "type": "string"
                        },
                        "SnapshotIdentifier": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VpcSecurityGroupIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "ClusterType",
                        "DBName",
                        "MasterUserPassword",
                        "MasterUsername",
                        "NodeType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Redshift::Cluster"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Redshift::Cluster.LoggingProperties": {
            "additionalProperties": false,
            "properties": {
                "BucketName": {
                    "type": "string"
                },
                "S3KeyPrefix": {
                    "type": "string"
                }
            },
            "required": [
                "BucketName"
            ],
            "type": "object"
        },
        "AWS::Redshift::ClusterParameterGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "ParameterGroupFamily": {
                            "type": "string"
                        },
                        "Parameters": {
                            "items": {
                                "$ref": "#/definitions/AWS::Redshift::ClusterParameterGroup.Parameter"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Description",
                        "ParameterGroupFamily"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Redshift::ClusterParameterGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Redshift::ClusterParameterGroup.Parameter": {
            "additionalProperties": false,
            "properties": {
                "ParameterName": {
                    "type": "string"
                },
                "ParameterValue": {
                    "type": "string"
                }
            },
            "required": [
                "ParameterName",
                "ParameterValue"
            ],
            "type": "object"
        },
        "AWS::Redshift::ClusterSecurityGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Description"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Redshift::ClusterSecurityGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Redshift::ClusterSecurityGroupIngress": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "CIDRIP": {
                            "type": "string"
                        },
                        "ClusterSecurityGroupName": {
                            "type": "string"
                        },
                        "EC2SecurityGroupName": {
                            "type": "string"
                        },
                        "EC2SecurityGroupOwnerId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ClusterSecurityGroupName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Redshift::ClusterSecurityGroupIngress"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Redshift::ClusterSubnetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "SubnetIds": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Description",
                        "SubnetIds"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Redshift::ClusterSubnetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Route53::HealthCheck": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "HealthCheckConfig": {
                            "$ref": "#/definitions/AWS::Route53::HealthCheck.HealthCheckConfig"
                        },
                        "HealthCheckTags": {
                            "items": {
                                "$ref": "#/definitions/AWS::Route53::HealthCheck.HealthCheckTag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "HealthCheckConfig"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Route53::HealthCheck"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Route53::HealthCheck.AlarmIdentifier": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Region": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Region"
            ],
            "type": "object"
        },
        "AWS::Route53::HealthCheck.HealthCheckConfig": {
            "additionalProperties": false,
            "properties": {
                "AlarmIdentifier": {
                    "$ref": "#/definitions/AWS::Route53::HealthCheck.AlarmIdentifier"
                },
                "ChildHealthChecks": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "EnableSNI": {
                    "type": "boolean"
                },
                "FailureThreshold": {
                    "type": "number"
                },
                "FullyQualifiedDomainName": {
                    "type": "string"
                },
                "HealthThreshold": {
                    "type": "number"
                },
                "IPAddress": {
                    "type": "string"
                },
                "InsufficientDataHealthStatus": {
                    "type": "string"
                },
                "Inverted": {
                    "type": "boolean"
                },
                "MeasureLatency": {
                    "type": "boolean"
                },
                "Port": {
                    "type": "number"
                },
                "Regions": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "RequestInterval": {
                    "type": "number"
                },
                "ResourcePath": {
                    "type": "string"
                },
                "SearchString": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Route53::HealthCheck.HealthCheckTag": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::Route53::HostedZone": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "HostedZoneConfig": {
                            "$ref": "#/definitions/AWS::Route53::HostedZone.HostedZoneConfig"
                        },
                        "HostedZoneTags": {
                            "items": {
                                "$ref": "#/definitions/AWS::Route53::HostedZone.HostedZoneTag"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "QueryLoggingConfig": {
                            "$ref": "#/definitions/AWS::Route53::HostedZone.QueryLoggingConfig"
                        },
                        "VPCs": {
                            "items": {
                                "$ref": "#/definitions/AWS::Route53::HostedZone.VPC"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Route53::HostedZone"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Route53::HostedZone.HostedZoneConfig": {
            "additionalProperties": false,
            "properties": {
                "Comment": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Route53::HostedZone.HostedZoneTag": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::Route53::HostedZone.QueryLoggingConfig": {
            "additionalProperties": false,
            "properties": {
                "CloudWatchLogsLogGroupArn": {
                    "type": "string"
                }
            },
            "required": [
                "CloudWatchLogsLogGroupArn"
            ],
            "type": "object"
        },
        "AWS::Route53::HostedZone.VPC": {
            "additionalProperties": false,
            "properties": {
                "VPCId": {
                    "type": "string"
                },
                "VPCRegion": {
                    "type": "string"
                }
            },
            "required": [
                "VPCId",
                "VPCRegion"
            ],
            "type": "object"
        },
        "AWS::Route53::RecordSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AliasTarget": {
                            "$ref": "#/definitions/AWS::Route53::RecordSet.AliasTarget"
                        },
                        "Comment": {
                            "type": "string"
                        },
                        "Failover": {
                            "type": "string"
                        },
                        "GeoLocation": {
                            "$ref": "#/definitions/AWS::Route53::RecordSet.GeoLocation"
                        },
                        "HealthCheckId": {
                            "type": "string"
                        },
                        "HostedZoneId": {
                            "type": "string"
                        },
                        "HostedZoneName": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Region": {
                            "type": "string"
                        },
                        "ResourceRecords": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "SetIdentifier": {
                            "type": "string"
                        },
                        "TTL": {
                            "type": "string"
                        },
                        "Type": {
                            "type": "string"
                        },
                        "Weight": {
                            "type": "number"
                        }
                    },
                    "required": [
                        "Name",
                        "Type"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Route53::RecordSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::Route53::RecordSet.AliasTarget": {
            "additionalProperties": false,
            "properties": {
                "DNSName": {
                    "type": "string"
                },
                "EvaluateTargetHealth": {
                    "type": "boolean"
                },
                "HostedZoneId": {
                    "type": "string"
                }
            },
            "required": [
                "DNSName",
                "HostedZoneId"
            ],
            "type": "object"
        },
        "AWS::Route53::RecordSet.GeoLocation": {
            "additionalProperties": false,
            "properties": {
                "ContinentCode": {
                    "type": "string"
                },
                "CountryCode": {
                    "type": "string"
                },
                "SubdivisionCode": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Route53::RecordSetGroup": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Comment": {
                            "type": "string"
                        },
                        "HostedZoneId": {
                            "type": "string"
                        },
                        "HostedZoneName": {
                            "type": "string"
                        },
                        "RecordSets": {
                            "items": {
                                "$ref": "#/definitions/AWS::Route53::RecordSetGroup.RecordSet"
                            },
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::Route53::RecordSetGroup"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::Route53::RecordSetGroup.AliasTarget": {
            "additionalProperties": false,
            "properties": {
                "DNSName": {
                    "type": "string"
                },
                "EvaluateTargetHealth": {
                    "type": "boolean"
                },
                "HostedZoneId": {
                    "type": "string"
                }
            },
            "required": [
                "DNSName",
                "HostedZoneId"
            ],
            "type": "object"
        },
        "AWS::Route53::RecordSetGroup.GeoLocation": {
            "additionalProperties": false,
            "properties": {
                "ContinentCode": {
                    "type": "string"
                },
                "CountryCode": {
                    "type": "string"
                },
                "SubdivisionCode": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::Route53::RecordSetGroup.RecordSet": {
            "additionalProperties": false,
            "properties": {
                "AliasTarget": {
                    "$ref": "#/definitions/AWS::Route53::RecordSetGroup.AliasTarget"
                },
                "Comment": {
                    "type": "string"
                },
                "Failover": {
                    "type": "string"
                },
                "GeoLocation": {
                    "$ref": "#/definitions/AWS::Route53::RecordSetGroup.GeoLocation"
                },
                "HealthCheckId": {
                    "type": "string"
                },
                "HostedZoneId": {
                    "type": "string"
                },
                "HostedZoneName": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Region": {
                    "type": "string"
                },
                "ResourceRecords": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "SetIdentifier": {
                    "type": "string"
                },
                "TTL": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                },
                "Weight": {
                    "type": "number"
                }
            },
            "required": [
                "Name",
                "Type"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AccelerateConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.AccelerateConfiguration"
                        },
                        "AccessControl": {
                            "type": "string"
                        },
                        "AnalyticsConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::S3::Bucket.AnalyticsConfiguration"
                            },
                            "type": "array"
                        },
                        "BucketEncryption": {
                            "$ref": "#/definitions/AWS::S3::Bucket.BucketEncryption"
                        },
                        "BucketName": {
                            "type": "string"
                        },
                        "CorsConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.CorsConfiguration"
                        },
                        "InventoryConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::S3::Bucket.InventoryConfiguration"
                            },
                            "type": "array"
                        },
                        "LifecycleConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.LifecycleConfiguration"
                        },
                        "LoggingConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.LoggingConfiguration"
                        },
                        "MetricsConfigurations": {
                            "items": {
                                "$ref": "#/definitions/AWS::S3::Bucket.MetricsConfiguration"
                            },
                            "type": "array"
                        },
                        "NotificationConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.NotificationConfiguration"
                        },
                        "ReplicationConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.ReplicationConfiguration"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        },
                        "VersioningConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.VersioningConfiguration"
                        },
                        "WebsiteConfiguration": {
                            "$ref": "#/definitions/AWS::S3::Bucket.WebsiteConfiguration"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::S3::Bucket"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.AbortIncompleteMultipartUpload": {
            "additionalProperties": false,
            "properties": {
                "DaysAfterInitiation": {
                    "type": "number"
                }
            },
            "required": [
                "DaysAfterInitiation"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.AccelerateConfiguration": {
            "additionalProperties": false,
            "properties": {
                "AccelerationStatus": {
                    "type": "string"
                }
            },
            "required": [
                "AccelerationStatus"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.AccessControlTranslation": {
            "additionalProperties": false,
            "properties": {
                "Owner": {
                    "type": "string"
                }
            },
            "required": [
                "Owner"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.AnalyticsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Id": {
                    "type": "string"
                },
                "Prefix": {
                    "type": "string"
                },
                "StorageClassAnalysis": {
                    "$ref": "#/definitions/AWS::S3::Bucket.StorageClassAnalysis"
                },
                "TagFilters": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.TagFilter"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Id",
                "StorageClassAnalysis"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.BucketEncryption": {
            "additionalProperties": false,
            "properties": {
                "ServerSideEncryptionConfiguration": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.ServerSideEncryptionRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "ServerSideEncryptionConfiguration"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.CorsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "CorsRules": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.CorsRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "CorsRules"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.CorsRule": {
            "additionalProperties": false,
            "properties": {
                "AllowedHeaders": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "AllowedMethods": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "AllowedOrigins": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "ExposedHeaders": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Id": {
                    "type": "string"
                },
                "MaxAge": {
                    "type": "number"
                }
            },
            "required": [
                "AllowedMethods",
                "AllowedOrigins"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.DataExport": {
            "additionalProperties": false,
            "properties": {
                "Destination": {
                    "$ref": "#/definitions/AWS::S3::Bucket.Destination"
                },
                "OutputSchemaVersion": {
                    "type": "string"
                }
            },
            "required": [
                "Destination",
                "OutputSchemaVersion"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.Destination": {
            "additionalProperties": false,
            "properties": {
                "BucketAccountId": {
                    "type": "string"
                },
                "BucketArn": {
                    "type": "string"
                },
                "Format": {
                    "type": "string"
                },
                "Prefix": {
                    "type": "string"
                }
            },
            "required": [
                "BucketArn",
                "Format"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.EncryptionConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ReplicaKmsKeyID": {
                    "type": "string"
                }
            },
            "required": [
                "ReplicaKmsKeyID"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.FilterRule": {
            "additionalProperties": false,
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Name",
                "Value"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.InventoryConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Destination": {
                    "$ref": "#/definitions/AWS::S3::Bucket.Destination"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "Id": {
                    "type": "string"
                },
                "IncludedObjectVersions": {
                    "type": "string"
                },
                "OptionalFields": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Prefix": {
                    "type": "string"
                },
                "ScheduleFrequency": {
                    "type": "string"
                }
            },
            "required": [
                "Destination",
                "Enabled",
                "Id",
                "IncludedObjectVersions",
                "ScheduleFrequency"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.LambdaConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Event": {
                    "type": "string"
                },
                "Filter": {
                    "$ref": "#/definitions/AWS::S3::Bucket.NotificationFilter"
                },
                "Function": {
                    "type": "string"
                }
            },
            "required": [
                "Event",
                "Function"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.LifecycleConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Rules": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.Rule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Rules"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.LoggingConfiguration": {
            "additionalProperties": false,
            "properties": {
                "DestinationBucketName": {
                    "type": "string"
                },
                "LogFilePrefix": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::S3::Bucket.MetricsConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Id": {
                    "type": "string"
                },
                "Prefix": {
                    "type": "string"
                },
                "TagFilters": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.TagFilter"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Id"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.NoncurrentVersionTransition": {
            "additionalProperties": false,
            "properties": {
                "StorageClass": {
                    "type": "string"
                },
                "TransitionInDays": {
                    "type": "number"
                }
            },
            "required": [
                "StorageClass",
                "TransitionInDays"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.NotificationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "LambdaConfigurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.LambdaConfiguration"
                    },
                    "type": "array"
                },
                "QueueConfigurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.QueueConfiguration"
                    },
                    "type": "array"
                },
                "TopicConfigurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.TopicConfiguration"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::S3::Bucket.NotificationFilter": {
            "additionalProperties": false,
            "properties": {
                "S3Key": {
                    "$ref": "#/definitions/AWS::S3::Bucket.S3KeyFilter"
                }
            },
            "required": [
                "S3Key"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.QueueConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Event": {
                    "type": "string"
                },
                "Filter": {
                    "$ref": "#/definitions/AWS::S3::Bucket.NotificationFilter"
                },
                "Queue": {
                    "type": "string"
                }
            },
            "required": [
                "Event",
                "Queue"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.RedirectAllRequestsTo": {
            "additionalProperties": false,
            "properties": {
                "HostName": {
                    "type": "string"
                },
                "Protocol": {
                    "type": "string"
                }
            },
            "required": [
                "HostName"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.RedirectRule": {
            "additionalProperties": false,
            "properties": {
                "HostName": {
                    "type": "string"
                },
                "HttpRedirectCode": {
                    "type": "string"
                },
                "Protocol": {
                    "type": "string"
                },
                "ReplaceKeyPrefixWith": {
                    "type": "string"
                },
                "ReplaceKeyWith": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::S3::Bucket.ReplicationConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Role": {
                    "type": "string"
                },
                "Rules": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.ReplicationRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Role",
                "Rules"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.ReplicationDestination": {
            "additionalProperties": false,
            "properties": {
                "AccessControlTranslation": {
                    "$ref": "#/definitions/AWS::S3::Bucket.AccessControlTranslation"
                },
                "Account": {
                    "type": "string"
                },
                "Bucket": {
                    "type": "string"
                },
                "EncryptionConfiguration": {
                    "$ref": "#/definitions/AWS::S3::Bucket.EncryptionConfiguration"
                },
                "StorageClass": {
                    "type": "string"
                }
            },
            "required": [
                "Bucket"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.ReplicationRule": {
            "additionalProperties": false,
            "properties": {
                "Destination": {
                    "$ref": "#/definitions/AWS::S3::Bucket.ReplicationDestination"
                },
                "Id": {
                    "type": "string"
                },
                "Prefix": {
                    "type": "string"
                },
                "SourceSelectionCriteria": {
                    "$ref": "#/definitions/AWS::S3::Bucket.SourceSelectionCriteria"
                },
                "Status": {
                    "type": "string"
                }
            },
            "required": [
                "Destination",
                "Prefix",
                "Status"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.RoutingRule": {
            "additionalProperties": false,
            "properties": {
                "RedirectRule": {
                    "$ref": "#/definitions/AWS::S3::Bucket.RedirectRule"
                },
                "RoutingRuleCondition": {
                    "$ref": "#/definitions/AWS::S3::Bucket.RoutingRuleCondition"
                }
            },
            "required": [
                "RedirectRule"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.RoutingRuleCondition": {
            "additionalProperties": false,
            "properties": {
                "HttpErrorCodeReturnedEquals": {
                    "type": "string"
                },
                "KeyPrefixEquals": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::S3::Bucket.Rule": {
            "additionalProperties": false,
            "properties": {
                "AbortIncompleteMultipartUpload": {
                    "$ref": "#/definitions/AWS::S3::Bucket.AbortIncompleteMultipartUpload"
                },
                "ExpirationDate": {
                    "type": "string"
                },
                "ExpirationInDays": {
                    "type": "number"
                },
                "Id": {
                    "type": "string"
                },
                "NoncurrentVersionExpirationInDays": {
                    "type": "number"
                },
                "NoncurrentVersionTransition": {
                    "$ref": "#/definitions/AWS::S3::Bucket.NoncurrentVersionTransition"
                },
                "NoncurrentVersionTransitions": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.NoncurrentVersionTransition"
                    },
                    "type": "array"
                },
                "Prefix": {
                    "type": "string"
                },
                "Status": {
                    "type": "string"
                },
                "TagFilters": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.TagFilter"
                    },
                    "type": "array"
                },
                "Transition": {
                    "$ref": "#/definitions/AWS::S3::Bucket.Transition"
                },
                "Transitions": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.Transition"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Status"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.S3KeyFilter": {
            "additionalProperties": false,
            "properties": {
                "Rules": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.FilterRule"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Rules"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.ServerSideEncryptionByDefault": {
            "additionalProperties": false,
            "properties": {
                "KMSMasterKeyID": {
                    "type": "string"
                },
                "SSEAlgorithm": {
                    "type": "string"
                }
            },
            "required": [
                "SSEAlgorithm"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.ServerSideEncryptionRule": {
            "additionalProperties": false,
            "properties": {
                "ServerSideEncryptionByDefault": {
                    "$ref": "#/definitions/AWS::S3::Bucket.ServerSideEncryptionByDefault"
                }
            },
            "type": "object"
        },
        "AWS::S3::Bucket.SourceSelectionCriteria": {
            "additionalProperties": false,
            "properties": {
                "SseKmsEncryptedObjects": {
                    "$ref": "#/definitions/AWS::S3::Bucket.SseKmsEncryptedObjects"
                }
            },
            "required": [
                "SseKmsEncryptedObjects"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.SseKmsEncryptedObjects": {
            "additionalProperties": false,
            "properties": {
                "Status": {
                    "type": "string"
                }
            },
            "required": [
                "Status"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.StorageClassAnalysis": {
            "additionalProperties": false,
            "properties": {
                "DataExport": {
                    "$ref": "#/definitions/AWS::S3::Bucket.DataExport"
                }
            },
            "type": "object"
        },
        "AWS::S3::Bucket.TagFilter": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Key",
                "Value"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.TopicConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Event": {
                    "type": "string"
                },
                "Filter": {
                    "$ref": "#/definitions/AWS::S3::Bucket.NotificationFilter"
                },
                "Topic": {
                    "type": "string"
                }
            },
            "required": [
                "Event",
                "Topic"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.Transition": {
            "additionalProperties": false,
            "properties": {
                "StorageClass": {
                    "type": "string"
                },
                "TransitionDate": {
                    "type": "string"
                },
                "TransitionInDays": {
                    "type": "number"
                }
            },
            "required": [
                "StorageClass"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.VersioningConfiguration": {
            "additionalProperties": false,
            "properties": {
                "Status": {
                    "type": "string"
                }
            },
            "required": [
                "Status"
            ],
            "type": "object"
        },
        "AWS::S3::Bucket.WebsiteConfiguration": {
            "additionalProperties": false,
            "properties": {
                "ErrorDocument": {
                    "type": "string"
                },
                "IndexDocument": {
                    "type": "string"
                },
                "RedirectAllRequestsTo": {
                    "$ref": "#/definitions/AWS::S3::Bucket.RedirectAllRequestsTo"
                },
                "RoutingRules": {
                    "items": {
                        "$ref": "#/definitions/AWS::S3::Bucket.RoutingRule"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::S3::BucketPolicy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Bucket": {
                            "type": "string"
                        },
                        "PolicyDocument": {
                            "type": "object"
                        }
                    },
                    "required": [
                        "Bucket",
                        "PolicyDocument"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::S3::BucketPolicy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SDB::Domain": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SDB::Domain"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SES::ConfigurationSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SES::ConfigurationSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SES::ConfigurationSetEventDestination": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ConfigurationSetName": {
                            "type": "string"
                        },
                        "EventDestination": {
                            "$ref": "#/definitions/AWS::SES::ConfigurationSetEventDestination.EventDestination"
                        }
                    },
                    "required": [
                        "ConfigurationSetName",
                        "EventDestination"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SES::ConfigurationSetEventDestination"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SES::ConfigurationSetEventDestination.CloudWatchDestination": {
            "additionalProperties": false,
            "properties": {
                "DimensionConfigurations": {
                    "items": {
                        "$ref": "#/definitions/AWS::SES::ConfigurationSetEventDestination.DimensionConfiguration"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::SES::ConfigurationSetEventDestination.DimensionConfiguration": {
            "additionalProperties": false,
            "properties": {
                "DefaultDimensionValue": {
                    "type": "string"
                },
                "DimensionName": {
                    "type": "string"
                },
                "DimensionValueSource": {
                    "type": "string"
                }
            },
            "required": [
                "DefaultDimensionValue",
                "DimensionName",
                "DimensionValueSource"
            ],
            "type": "object"
        },
        "AWS::SES::ConfigurationSetEventDestination.EventDestination": {
            "additionalProperties": false,
            "properties": {
                "CloudWatchDestination": {
                    "$ref": "#/definitions/AWS::SES::ConfigurationSetEventDestination.CloudWatchDestination"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "KinesisFirehoseDestination": {
                    "$ref": "#/definitions/AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination"
                },
                "MatchingEventTypes": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "MatchingEventTypes"
            ],
            "type": "object"
        },
        "AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination": {
            "additionalProperties": false,
            "properties": {
                "DeliveryStreamARN": {
                    "type": "string"
                },
                "IAMRoleARN": {
                    "type": "string"
                }
            },
            "required": [
                "DeliveryStreamARN",
                "IAMRoleARN"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptFilter": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Filter": {
                            "$ref": "#/definitions/AWS::SES::ReceiptFilter.Filter"
                        }
                    },
                    "required": [
                        "Filter"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SES::ReceiptFilter"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptFilter.Filter": {
            "additionalProperties": false,
            "properties": {
                "IpFilter": {
                    "$ref": "#/definitions/AWS::SES::ReceiptFilter.IpFilter"
                },
                "Name": {
                    "type": "string"
                }
            },
            "required": [
                "IpFilter"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptFilter.IpFilter": {
            "additionalProperties": false,
            "properties": {
                "Cidr": {
                    "type": "string"
                },
                "Policy": {
                    "type": "string"
                }
            },
            "required": [
                "Cidr",
                "Policy"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "After": {
                            "type": "string"
                        },
                        "Rule": {
                            "$ref": "#/definitions/AWS::SES::ReceiptRule.Rule"
                        },
                        "RuleSetName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Rule",
                        "RuleSetName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SES::ReceiptRule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule.Action": {
            "additionalProperties": false,
            "properties": {
                "AddHeaderAction": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.AddHeaderAction"
                },
                "BounceAction": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.BounceAction"
                },
                "LambdaAction": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.LambdaAction"
                },
                "S3Action": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.S3Action"
                },
                "SNSAction": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.SNSAction"
                },
                "StopAction": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.StopAction"
                },
                "WorkmailAction": {
                    "$ref": "#/definitions/AWS::SES::ReceiptRule.WorkmailAction"
                }
            },
            "type": "object"
        },
        "AWS::SES::ReceiptRule.AddHeaderAction": {
            "additionalProperties": false,
            "properties": {
                "HeaderName": {
                    "type": "string"
                },
                "HeaderValue": {
                    "type": "string"
                }
            },
            "required": [
                "HeaderName",
                "HeaderValue"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule.BounceAction": {
            "additionalProperties": false,
            "properties": {
                "Message": {
                    "type": "string"
                },
                "Sender": {
                    "type": "string"
                },
                "SmtpReplyCode": {
                    "type": "string"
                },
                "StatusCode": {
                    "type": "string"
                },
                "TopicArn": {
                    "type": "string"
                }
            },
            "required": [
                "Message",
                "Sender",
                "SmtpReplyCode"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule.LambdaAction": {
            "additionalProperties": false,
            "properties": {
                "FunctionArn": {
                    "type": "string"
                },
                "InvocationType": {
                    "type": "string"
                },
                "TopicArn": {
                    "type": "string"
                }
            },
            "required": [
                "FunctionArn"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule.Rule": {
            "additionalProperties": false,
            "properties": {
                "Actions": {
                    "items": {
                        "$ref": "#/definitions/AWS::SES::ReceiptRule.Action"
                    },
                    "type": "array"
                },
                "Enabled": {
                    "type": "boolean"
                },
                "Name": {
                    "type": "string"
                },
                "Recipients": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "ScanEnabled": {
                    "type": "boolean"
                },
                "TlsPolicy": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::SES::ReceiptRule.S3Action": {
            "additionalProperties": false,
            "properties": {
                "BucketName": {
                    "type": "string"
                },
                "KmsKeyArn": {
                    "type": "string"
                },
                "ObjectKeyPrefix": {
                    "type": "string"
                },
                "TopicArn": {
                    "type": "string"
                }
            },
            "required": [
                "BucketName"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule.SNSAction": {
            "additionalProperties": false,
            "properties": {
                "Encoding": {
                    "type": "string"
                },
                "TopicArn": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::SES::ReceiptRule.StopAction": {
            "additionalProperties": false,
            "properties": {
                "Scope": {
                    "type": "string"
                },
                "TopicArn": {
                    "type": "string"
                }
            },
            "required": [
                "Scope"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRule.WorkmailAction": {
            "additionalProperties": false,
            "properties": {
                "OrganizationArn": {
                    "type": "string"
                },
                "TopicArn": {
                    "type": "string"
                }
            },
            "required": [
                "OrganizationArn"
            ],
            "type": "object"
        },
        "AWS::SES::ReceiptRuleSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "RuleSetName": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SES::ReceiptRuleSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SES::Template": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Template": {
                            "$ref": "#/definitions/AWS::SES::Template.Template"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SES::Template"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SES::Template.Template": {
            "additionalProperties": false,
            "properties": {
                "HtmlPart": {
                    "type": "string"
                },
                "SubjectPart": {
                    "type": "string"
                },
                "TemplateName": {
                    "type": "string"
                },
                "TextPart": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::SNS::Subscription": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Endpoint": {
                            "type": "string"
                        },
                        "Protocol": {
                            "type": "string"
                        },
                        "TopicArn": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SNS::Subscription"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SNS::Topic": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DisplayName": {
                            "type": "string"
                        },
                        "Subscription": {
                            "items": {
                                "$ref": "#/definitions/AWS::SNS::Topic.Subscription"
                            },
                            "type": "array"
                        },
                        "TopicName": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SNS::Topic"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SNS::Topic.Subscription": {
            "additionalProperties": false,
            "properties": {
                "Endpoint": {
                    "type": "string"
                },
                "Protocol": {
                    "type": "string"
                }
            },
            "required": [
                "Endpoint",
                "Protocol"
            ],
            "type": "object"
        },
        "AWS::SNS::TopicPolicy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PolicyDocument": {
                            "type": "object"
                        },
                        "Topics": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "PolicyDocument",
                        "Topics"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SNS::TopicPolicy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SQS::Queue": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ContentBasedDeduplication": {
                            "type": "boolean"
                        },
                        "DelaySeconds": {
                            "type": "number"
                        },
                        "FifoQueue": {
                            "type": "boolean"
                        },
                        "KmsDataKeyReusePeriodSeconds": {
                            "type": "number"
                        },
                        "KmsMasterKeyId": {
                            "type": "string"
                        },
                        "MaximumMessageSize": {
                            "type": "number"
                        },
                        "MessageRetentionPeriod": {
                            "type": "number"
                        },
                        "QueueName": {
                            "type": "string"
                        },
                        "ReceiveMessageWaitTimeSeconds": {
                            "type": "number"
                        },
                        "RedrivePolicy": {
                            "type": "object"
                        },
                        "VisibilityTimeout": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SQS::Queue"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::SQS::QueuePolicy": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "PolicyDocument": {
                            "type": "object"
                        },
                        "Queues": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "PolicyDocument",
                        "Queues"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SQS::QueuePolicy"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SSM::Association": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AssociationName": {
                            "type": "string"
                        },
                        "DocumentVersion": {
                            "type": "string"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Parameters": {
                            "additionalProperties": false,
                            "patternProperties": {
                                "^[a-zA-Z0-9]+$": {
                                    "$ref": "#/definitions/AWS::SSM::Association.ParameterValues"
                                }
                            },
                            "type": "object"
                        },
                        "ScheduleExpression": {
                            "type": "string"
                        },
                        "Targets": {
                            "items": {
                                "$ref": "#/definitions/AWS::SSM::Association.Target"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SSM::Association"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SSM::Association.ParameterValues": {
            "additionalProperties": false,
            "properties": {
                "ParameterValues": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "ParameterValues"
            ],
            "type": "object"
        },
        "AWS::SSM::Association.Target": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Key",
                "Values"
            ],
            "type": "object"
        },
        "AWS::SSM::Document": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Content": {
                            "type": "object"
                        },
                        "DocumentType": {
                            "type": "string"
                        },
                        "Tags": {
                            "items": {
                                "$ref": "#/definitions/Tag"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Content"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SSM::Document"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "LoggingInfo": {
                            "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.LoggingInfo"
                        },
                        "MaxConcurrency": {
                            "type": "string"
                        },
                        "MaxErrors": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Priority": {
                            "type": "number"
                        },
                        "ServiceRoleArn": {
                            "type": "string"
                        },
                        "Targets": {
                            "items": {
                                "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.Target"
                            },
                            "type": "array"
                        },
                        "TaskArn": {
                            "type": "string"
                        },
                        "TaskInvocationParameters": {
                            "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.TaskInvocationParameters"
                        },
                        "TaskParameters": {
                            "type": "object"
                        },
                        "TaskType": {
                            "type": "string"
                        },
                        "WindowId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "MaxConcurrency",
                        "MaxErrors",
                        "Priority",
                        "ServiceRoleArn",
                        "Targets",
                        "TaskArn",
                        "TaskType"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SSM::MaintenanceWindowTask"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.LoggingInfo": {
            "additionalProperties": false,
            "properties": {
                "Region": {
                    "type": "string"
                },
                "S3Bucket": {
                    "type": "string"
                },
                "S3Prefix": {
                    "type": "string"
                }
            },
            "required": [
                "Region",
                "S3Bucket"
            ],
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowAutomationParameters": {
            "additionalProperties": false,
            "properties": {
                "DocumentVersion": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                }
            },
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowLambdaParameters": {
            "additionalProperties": false,
            "properties": {
                "ClientContext": {
                    "type": "string"
                },
                "Payload": {
                    "type": "string"
                },
                "Qualifier": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowRunCommandParameters": {
            "additionalProperties": false,
            "properties": {
                "Comment": {
                    "type": "string"
                },
                "DocumentHash": {
                    "type": "string"
                },
                "DocumentHashType": {
                    "type": "string"
                },
                "NotificationConfig": {
                    "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.NotificationConfig"
                },
                "OutputS3BucketName": {
                    "type": "string"
                },
                "OutputS3KeyPrefix": {
                    "type": "string"
                },
                "Parameters": {
                    "type": "object"
                },
                "ServiceRoleArn": {
                    "type": "string"
                },
                "TimeoutSeconds": {
                    "type": "number"
                }
            },
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowStepFunctionsParameters": {
            "additionalProperties": false,
            "properties": {
                "Input": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                }
            },
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.NotificationConfig": {
            "additionalProperties": false,
            "properties": {
                "NotificationArn": {
                    "type": "string"
                },
                "NotificationEvents": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "NotificationType": {
                    "type": "string"
                }
            },
            "required": [
                "NotificationArn"
            ],
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.Target": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "required": [
                "Key"
            ],
            "type": "object"
        },
        "AWS::SSM::MaintenanceWindowTask.TaskInvocationParameters": {
            "additionalProperties": false,
            "properties": {
                "MaintenanceWindowAutomationParameters": {
                    "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.MaintenanceWindowAutomationParameters"
                },
                "MaintenanceWindowLambdaParameters": {
                    "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.MaintenanceWindowLambdaParameters"
                },
                "MaintenanceWindowRunCommandParameters": {
                    "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.MaintenanceWindowRunCommandParameters"
                },
                "MaintenanceWindowStepFunctionsParameters": {
                    "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask.MaintenanceWindowStepFunctionsParameters"
                }
            },
            "type": "object"
        },
        "AWS::SSM::Parameter": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "AllowedPattern": {
                            "type": "string"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Type": {
                            "type": "string"
                        },
                        "Value": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Type",
                        "Value"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SSM::Parameter"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SSM::PatchBaseline": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ApprovalRules": {
                            "$ref": "#/definitions/AWS::SSM::PatchBaseline.RuleGroup"
                        },
                        "ApprovedPatches": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "ApprovedPatchesComplianceLevel": {
                            "type": "string"
                        },
                        "ApprovedPatchesEnableNonSecurity": {
                            "type": "boolean"
                        },
                        "Description": {
                            "type": "string"
                        },
                        "GlobalFilters": {
                            "$ref": "#/definitions/AWS::SSM::PatchBaseline.PatchFilterGroup"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "OperatingSystem": {
                            "type": "string"
                        },
                        "PatchGroups": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "RejectedPatches": {
                            "items": {
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "Sources": {
                            "items": {
                                "$ref": "#/definitions/AWS::SSM::PatchBaseline.PatchSource"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::SSM::PatchBaseline"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::SSM::PatchBaseline.PatchFilter": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Values": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::SSM::PatchBaseline.PatchFilterGroup": {
            "additionalProperties": false,
            "properties": {
                "PatchFilters": {
                    "items": {
                        "$ref": "#/definitions/AWS::SSM::PatchBaseline.PatchFilter"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::SSM::PatchBaseline.PatchSource": {
            "additionalProperties": false,
            "properties": {
                "Configuration": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Products": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::SSM::PatchBaseline.Rule": {
            "additionalProperties": false,
            "properties": {
                "ApproveAfterDays": {
                    "type": "number"
                },
                "ComplianceLevel": {
                    "type": "string"
                },
                "EnableNonSecurity": {
                    "type": "boolean"
                },
                "PatchFilterGroup": {
                    "$ref": "#/definitions/AWS::SSM::PatchBaseline.PatchFilterGroup"
                }
            },
            "type": "object"
        },
        "AWS::SSM::PatchBaseline.RuleGroup": {
            "additionalProperties": false,
            "properties": {
                "PatchRules": {
                    "items": {
                        "$ref": "#/definitions/AWS::SSM::PatchBaseline.Rule"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        },
        "AWS::ServiceDiscovery::Instance": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "InstanceAttributes": {
                            "type": "object"
                        },
                        "InstanceId": {
                            "type": "string"
                        },
                        "ServiceId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "InstanceAttributes",
                        "ServiceId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ServiceDiscovery::Instance"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ServiceDiscovery::PrivateDnsNamespace": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Vpc": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name",
                        "Vpc"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ServiceDiscovery::PrivateDnsNamespace"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ServiceDiscovery::PublicDnsNamespace": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ServiceDiscovery::PublicDnsNamespace"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ServiceDiscovery::Service": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Description": {
                            "type": "string"
                        },
                        "DnsConfig": {
                            "$ref": "#/definitions/AWS::ServiceDiscovery::Service.DnsConfig"
                        },
                        "HealthCheckConfig": {
                            "$ref": "#/definitions/AWS::ServiceDiscovery::Service.HealthCheckConfig"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DnsConfig"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::ServiceDiscovery::Service"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::ServiceDiscovery::Service.DnsConfig": {
            "additionalProperties": false,
            "properties": {
                "DnsRecords": {
                    "items": {
                        "$ref": "#/definitions/AWS::ServiceDiscovery::Service.DnsRecord"
                    },
                    "type": "array"
                },
                "NamespaceId": {
                    "type": "string"
                }
            },
            "required": [
                "DnsRecords",
                "NamespaceId"
            ],
            "type": "object"
        },
        "AWS::ServiceDiscovery::Service.DnsRecord": {
            "additionalProperties": false,
            "properties": {
                "TTL": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "TTL",
                "Type"
            ],
            "type": "object"
        },
        "AWS::ServiceDiscovery::Service.HealthCheckConfig": {
            "additionalProperties": false,
            "properties": {
                "FailureThreshold": {
                    "type": "number"
                },
                "ResourcePath": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::StepFunctions::Activity": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::StepFunctions::Activity"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::StepFunctions::StateMachine": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DefinitionString": {
                            "type": "string"
                        },
                        "RoleArn": {
                            "type": "string"
                        },
                        "StateMachineName": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "DefinitionString",
                        "RoleArn"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::StepFunctions::StateMachine"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::ByteMatchSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ByteMatchTuples": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::ByteMatchSet.ByteMatchTuple"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::ByteMatchSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::ByteMatchSet.ByteMatchTuple": {
            "additionalProperties": false,
            "properties": {
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAF::ByteMatchSet.FieldToMatch"
                },
                "PositionalConstraint": {
                    "type": "string"
                },
                "TargetString": {
                    "type": "string"
                },
                "TargetStringBase64": {
                    "type": "string"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "FieldToMatch",
                "PositionalConstraint",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAF::ByteMatchSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAF::IPSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "IPSetDescriptors": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::IPSet.IPSetDescriptor"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::IPSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::IPSet.IPSetDescriptor": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Value"
            ],
            "type": "object"
        },
        "AWS::WAF::Rule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "MetricName": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Predicates": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::Rule.Predicate"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "MetricName",
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::Rule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::Rule.Predicate": {
            "additionalProperties": false,
            "properties": {
                "DataId": {
                    "type": "string"
                },
                "Negated": {
                    "type": "boolean"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "DataId",
                "Negated",
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAF::SizeConstraintSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "SizeConstraints": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::SizeConstraintSet.SizeConstraint"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name",
                        "SizeConstraints"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::SizeConstraintSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::SizeConstraintSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAF::SizeConstraintSet.SizeConstraint": {
            "additionalProperties": false,
            "properties": {
                "ComparisonOperator": {
                    "type": "string"
                },
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAF::SizeConstraintSet.FieldToMatch"
                },
                "Size": {
                    "type": "number"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "ComparisonOperator",
                "FieldToMatch",
                "Size",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAF::SqlInjectionMatchSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "SqlInjectionMatchTuples": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::SqlInjectionMatchSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::SqlInjectionMatchSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple": {
            "additionalProperties": false,
            "properties": {
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAF::SqlInjectionMatchSet.FieldToMatch"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "FieldToMatch",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAF::WebACL": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DefaultAction": {
                            "$ref": "#/definitions/AWS::WAF::WebACL.WafAction"
                        },
                        "MetricName": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Rules": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::WebACL.ActivatedRule"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DefaultAction",
                        "MetricName",
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::WebACL"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::WebACL.ActivatedRule": {
            "additionalProperties": false,
            "properties": {
                "Action": {
                    "$ref": "#/definitions/AWS::WAF::WebACL.WafAction"
                },
                "Priority": {
                    "type": "number"
                },
                "RuleId": {
                    "type": "string"
                }
            },
            "required": [
                "Action",
                "Priority",
                "RuleId"
            ],
            "type": "object"
        },
        "AWS::WAF::WebACL.WafAction": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAF::XssMatchSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "XssMatchTuples": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAF::XssMatchSet.XssMatchTuple"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name",
                        "XssMatchTuples"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAF::XssMatchSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAF::XssMatchSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAF::XssMatchSet.XssMatchTuple": {
            "additionalProperties": false,
            "properties": {
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAF::XssMatchSet.FieldToMatch"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "FieldToMatch",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::ByteMatchSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ByteMatchTuples": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::ByteMatchSet.ByteMatchTuple"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::ByteMatchSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::ByteMatchSet.ByteMatchTuple": {
            "additionalProperties": false,
            "properties": {
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAFRegional::ByteMatchSet.FieldToMatch"
                },
                "PositionalConstraint": {
                    "type": "string"
                },
                "TargetString": {
                    "type": "string"
                },
                "TargetStringBase64": {
                    "type": "string"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "FieldToMatch",
                "PositionalConstraint",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::ByteMatchSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::IPSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "IPSetDescriptors": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::IPSet.IPSetDescriptor"
                            },
                            "type": "array"
                        },
                        "Name": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::IPSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::IPSet.IPSetDescriptor": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Value"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::Rule": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "MetricName": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Predicates": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::Rule.Predicate"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "MetricName",
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::Rule"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::Rule.Predicate": {
            "additionalProperties": false,
            "properties": {
                "DataId": {
                    "type": "string"
                },
                "Negated": {
                    "type": "boolean"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "DataId",
                "Negated",
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::SizeConstraintSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "SizeConstraints": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::SizeConstraintSet.SizeConstraint"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::SizeConstraintSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::SizeConstraintSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::SizeConstraintSet.SizeConstraint": {
            "additionalProperties": false,
            "properties": {
                "ComparisonOperator": {
                    "type": "string"
                },
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAFRegional::SizeConstraintSet.FieldToMatch"
                },
                "Size": {
                    "type": "number"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "ComparisonOperator",
                "FieldToMatch",
                "Size",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::SqlInjectionMatchSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "SqlInjectionMatchTuples": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::SqlInjectionMatchSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::SqlInjectionMatchSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple": {
            "additionalProperties": false,
            "properties": {
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAFRegional::SqlInjectionMatchSet.FieldToMatch"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "FieldToMatch",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::WebACL": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "DefaultAction": {
                            "$ref": "#/definitions/AWS::WAFRegional::WebACL.Action"
                        },
                        "MetricName": {
                            "type": "string"
                        },
                        "Name": {
                            "type": "string"
                        },
                        "Rules": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::WebACL.Rule"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "DefaultAction",
                        "MetricName",
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::WebACL"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::WebACL.Action": {
            "additionalProperties": false,
            "properties": {
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::WebACL.Rule": {
            "additionalProperties": false,
            "properties": {
                "Action": {
                    "$ref": "#/definitions/AWS::WAFRegional::WebACL.Action"
                },
                "Priority": {
                    "type": "number"
                },
                "RuleId": {
                    "type": "string"
                }
            },
            "required": [
                "Action",
                "Priority",
                "RuleId"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::WebACLAssociation": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "ResourceArn": {
                            "type": "string"
                        },
                        "WebACLId": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "ResourceArn",
                        "WebACLId"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::WebACLAssociation"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::XssMatchSet": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "Name": {
                            "type": "string"
                        },
                        "XssMatchTuples": {
                            "items": {
                                "$ref": "#/definitions/AWS::WAFRegional::XssMatchSet.XssMatchTuple"
                            },
                            "type": "array"
                        }
                    },
                    "required": [
                        "Name"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WAFRegional::XssMatchSet"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::XssMatchSet.FieldToMatch": {
            "additionalProperties": false,
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "AWS::WAFRegional::XssMatchSet.XssMatchTuple": {
            "additionalProperties": false,
            "properties": {
                "FieldToMatch": {
                    "$ref": "#/definitions/AWS::WAFRegional::XssMatchSet.FieldToMatch"
                },
                "TextTransformation": {
                    "type": "string"
                }
            },
            "required": [
                "FieldToMatch",
                "TextTransformation"
            ],
            "type": "object"
        },
        "AWS::WorkSpaces::Workspace": {
            "additionalProperties": false,
            "properties": {
                "DeletionPolicy": {
                    "enum": [
                        "Delete",
                        "Retain",
                        "Snapshot"
                    ],
                    "type": "string"
                },
                "DependsOn": {
                    "anyOf": [
                        {
                            "pattern": "^[a-zA-Z0-9]+$",
                            "type": "string"
                        },
                        {
                            "items": {
                                "pattern": "^[a-zA-Z0-9]+$",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    ]
                },
                "Metadata": {
                    "type": "object"
                },
                "Properties": {
                    "additionalProperties": false,
                    "properties": {
                        "BundleId": {
                            "type": "string"
                        },
                        "DirectoryId": {
                            "type": "string"
                        },
                        "RootVolumeEncryptionEnabled": {
                            "type": "boolean"
                        },
                        "UserName": {
                            "type": "string"
                        },
                        "UserVolumeEncryptionEnabled": {
                            "type": "boolean"
                        },
                        "VolumeEncryptionKey": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "BundleId",
                        "DirectoryId",
                        "UserName"
                    ],
                    "type": "object"
                },
                "Type": {
                    "enum": [
                        "AWS::WorkSpaces::Workspace"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type",
                "Properties"
            ],
            "type": "object"
        },
        "Parameter": {
            "additionalProperties": false,
            "properties": {
                "AllowedPattern": {
                    "type": "string"
                },
                "AllowedValues": {
                    "type": "array"
                },
                "ConstraintDescription": {
                    "type": "string"
                },
                "Default": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "MaxLength": {
                    "type": "string"
                },
                "MaxValue": {
                    "type": "string"
                },
                "MinLength": {
                    "type": "string"
                },
                "MinValue": {
                    "type": "string"
                },
                "NoEcho": {
                    "type": [
                        "string",
                        "boolean"
                    ]
                },
                "Type": {
                    "enum": [
                        "String",
                        "Number",
                        "List\u003cNumber\u003e",
                        "CommaDelimitedList",
                        "AWS::EC2::AvailabilityZone::Name",
                        "AWS::EC2::Image::Id",
                        "AWS::EC2::Instance::Id",
                        "AWS::EC2::KeyPair::KeyName",
                        "AWS::EC2::SecurityGroup::GroupName",
                        "AWS::EC2::SecurityGroup::Id",
                        "AWS::EC2::Subnet::Id",
                        "AWS::EC2::Volume::Id",
                        "AWS::EC2::VPC::Id",
                        "AWS::Route53::HostedZone::Id",
                        "List\u003cAWS::EC2::AvailabilityZone::Name\u003e",
                        "List\u003cAWS::EC2::Image::Id\u003e",
                        "List\u003cAWS::EC2::Instance::Id\u003e",
                        "List\u003cAWS::EC2::SecurityGroup::GroupName\u003e",
                        "List\u003cAWS::EC2::SecurityGroup::Id\u003e",
                        "List\u003cAWS::EC2::Subnet::Id\u003e",
                        "List\u003cAWS::EC2::Volume::Id\u003e",
                        "List\u003cAWS::EC2::VPC::Id\u003e",
                        "List\u003cAWS::Route53::HostedZone::Id\u003e",
                        "List\u003cString\u003e"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "Type"
            ],
            "type": "object"
        },
        "Tag": {
            "additionalProperties": false,
            "properties": {
                "Key": {
                    "type": "string"
                },
                "Value": {
                    "type": "string"
                }
            },
            "type": "object"
        }
    },
    "properties": {
        "AWSTemplateFormatVersion": {
            "enum": [
                "2010-09-09"
            ],
            "type": "string"
        },
        "Conditions": {
            "additionalProperties": false,
            "patternProperties": {
                "^[a-zA-Z0-9]+$": {
                    "type": "object"
                }
            },
            "type": "object"
        },
        "Description": {
            "description": "Template description",
            "maxLength": 1024,
            "type": "string"
        },
        "Mappings": {
            "additionalProperties": false,
            "patternProperties": {
                "^[a-zA-Z0-9]+$": {
                    "type": "object"
                }
            },
            "type": "object"
        },
        "Metadata": {
            "type": "object"
        },
        "Outputs": {
            "additionalProperties": false,
            "maxProperties": 60,
            "minProperties": 1,
            "patternProperties": {
                "^[a-zA-Z0-9]+$": {
                    "type": "object"
                }
            },
            "type": "object"
        },
        "Parameters": {
            "additionalProperties": false,
            "maxProperties": 50,
            "patternProperties": {
                "^[a-zA-Z0-9]+$": {
                    "$ref": "#/definitions/Parameter"
                }
            },
            "type": "object"
        },
        "Resources": {
            "additionalProperties": false,
            "patternProperties": {
                "^[a-zA-Z0-9]+$": {
                    "anyOf": [
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Account"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::ApiKey"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Authorizer"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::BasePathMapping"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::ClientCertificate"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Deployment"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::DocumentationPart"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::DocumentationVersion"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::DomainName"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::GatewayResponse"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Method"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Model"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::RequestValidator"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Resource"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::RestApi"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::Stage"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::UsagePlan"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::UsagePlanKey"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApiGateway::VpcLink"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalableTarget"
                        },
                        {
                            "$ref": "#/definitions/AWS::ApplicationAutoScaling::ScalingPolicy"
                        },
                        {
                            "$ref": "#/definitions/AWS::Athena::NamedQuery"
                        },
                        {
                            "$ref": "#/definitions/AWS::AutoScaling::AutoScalingGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::AutoScaling::LaunchConfiguration"
                        },
                        {
                            "$ref": "#/definitions/AWS::AutoScaling::LifecycleHook"
                        },
                        {
                            "$ref": "#/definitions/AWS::AutoScaling::ScalingPolicy"
                        },
                        {
                            "$ref": "#/definitions/AWS::AutoScaling::ScheduledAction"
                        },
                        {
                            "$ref": "#/definitions/AWS::Batch::ComputeEnvironment"
                        },
                        {
                            "$ref": "#/definitions/AWS::Batch::JobDefinition"
                        },
                        {
                            "$ref": "#/definitions/AWS::Batch::JobQueue"
                        },
                        {
                            "$ref": "#/definitions/AWS::CertificateManager::Certificate"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cloud9::EnvironmentEC2"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFormation::CustomResource"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFormation::Stack"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFormation::WaitCondition"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFormation::WaitConditionHandle"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFront::CloudFrontOriginAccessIdentity"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFront::Distribution"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudFront::StreamingDistribution"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudTrail::Trail"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudWatch::Alarm"
                        },
                        {
                            "$ref": "#/definitions/AWS::CloudWatch::Dashboard"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodeBuild::Project"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodeCommit::Repository"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodeDeploy::Application"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentConfig"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodeDeploy::DeploymentGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodePipeline::CustomActionType"
                        },
                        {
                            "$ref": "#/definitions/AWS::CodePipeline::Pipeline"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::IdentityPool"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::IdentityPoolRoleAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::UserPool"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::UserPoolClient"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::UserPoolGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::UserPoolUser"
                        },
                        {
                            "$ref": "#/definitions/AWS::Cognito::UserPoolUserToGroupAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::Config::ConfigRule"
                        },
                        {
                            "$ref": "#/definitions/AWS::Config::ConfigurationRecorder"
                        },
                        {
                            "$ref": "#/definitions/AWS::Config::DeliveryChannel"
                        },
                        {
                            "$ref": "#/definitions/AWS::DAX::Cluster"
                        },
                        {
                            "$ref": "#/definitions/AWS::DAX::ParameterGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::DAX::SubnetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::DMS::Certificate"
                        },
                        {
                            "$ref": "#/definitions/AWS::DMS::Endpoint"
                        },
                        {
                            "$ref": "#/definitions/AWS::DMS::EventSubscription"
                        },
                        {
                            "$ref": "#/definitions/AWS::DMS::ReplicationInstance"
                        },
                        {
                            "$ref": "#/definitions/AWS::DMS::ReplicationSubnetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::DMS::ReplicationTask"
                        },
                        {
                            "$ref": "#/definitions/AWS::DataPipeline::Pipeline"
                        },
                        {
                            "$ref": "#/definitions/AWS::DirectoryService::MicrosoftAD"
                        },
                        {
                            "$ref": "#/definitions/AWS::DirectoryService::SimpleAD"
                        },
                        {
                            "$ref": "#/definitions/AWS::DynamoDB::Table"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::CustomerGateway"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::DHCPOptions"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::EIP"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::EIPAssociation"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::EgressOnlyInternetGateway"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::FlowLog"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::Host"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::Instance"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::InternetGateway"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::NatGateway"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::NetworkAcl"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::NetworkAclEntry"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::NetworkInterface"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::NetworkInterfaceAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::NetworkInterfacePermission"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::PlacementGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::Route"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::RouteTable"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SecurityGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SecurityGroupEgress"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SecurityGroupIngress"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SpotFleet"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::Subnet"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SubnetCidrBlock"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SubnetNetworkAclAssociation"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::SubnetRouteTableAssociation"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::TrunkInterfaceAssociation"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPC"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPCCidrBlock"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPCDHCPOptionsAssociation"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPCEndpoint"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPCGatewayAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPCPeeringConnection"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPNConnection"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPNConnectionRoute"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPNGateway"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VPNGatewayRoutePropagation"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::Volume"
                        },
                        {
                            "$ref": "#/definitions/AWS::EC2::VolumeAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::ECR::Repository"
                        },
                        {
                            "$ref": "#/definitions/AWS::ECS::Cluster"
                        },
                        {
                            "$ref": "#/definitions/AWS::ECS::Service"
                        },
                        {
                            "$ref": "#/definitions/AWS::ECS::TaskDefinition"
                        },
                        {
                            "$ref": "#/definitions/AWS::EFS::FileSystem"
                        },
                        {
                            "$ref": "#/definitions/AWS::EFS::MountTarget"
                        },
                        {
                            "$ref": "#/definitions/AWS::EMR::Cluster"
                        },
                        {
                            "$ref": "#/definitions/AWS::EMR::InstanceFleetConfig"
                        },
                        {
                            "$ref": "#/definitions/AWS::EMR::InstanceGroupConfig"
                        },
                        {
                            "$ref": "#/definitions/AWS::EMR::SecurityConfiguration"
                        },
                        {
                            "$ref": "#/definitions/AWS::EMR::Step"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElastiCache::CacheCluster"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElastiCache::ParameterGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElastiCache::ReplicationGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElastiCache::SecurityGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElastiCache::SecurityGroupIngress"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElastiCache::SubnetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::Application"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::ApplicationVersion"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::ConfigurationTemplate"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticBeanstalk::Environment"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancing::LoadBalancer"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::Listener"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::ListenerCertificate"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::ListenerRule"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::LoadBalancer"
                        },
                        {
                            "$ref": "#/definitions/AWS::ElasticLoadBalancingV2::TargetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Elasticsearch::Domain"
                        },
                        {
                            "$ref": "#/definitions/AWS::Events::Rule"
                        },
                        {
                            "$ref": "#/definitions/AWS::GameLift::Alias"
                        },
                        {
                            "$ref": "#/definitions/AWS::GameLift::Build"
                        },
                        {
                            "$ref": "#/definitions/AWS::GameLift::Fleet"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Classifier"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Connection"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Crawler"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Database"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::DevEndpoint"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Job"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Partition"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Table"
                        },
                        {
                            "$ref": "#/definitions/AWS::Glue::Trigger"
                        },
                        {
                            "$ref": "#/definitions/AWS::GuardDuty::Detector"
                        },
                        {
                            "$ref": "#/definitions/AWS::GuardDuty::IPSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::GuardDuty::Master"
                        },
                        {
                            "$ref": "#/definitions/AWS::GuardDuty::Member"
                        },
                        {
                            "$ref": "#/definitions/AWS::GuardDuty::ThreatIntelSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::AccessKey"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::Group"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::InstanceProfile"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::ManagedPolicy"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::Policy"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::Role"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::User"
                        },
                        {
                            "$ref": "#/definitions/AWS::IAM::UserToGroupAddition"
                        },
                        {
                            "$ref": "#/definitions/AWS::Inspector::AssessmentTarget"
                        },
                        {
                            "$ref": "#/definitions/AWS::Inspector::AssessmentTemplate"
                        },
                        {
                            "$ref": "#/definitions/AWS::Inspector::ResourceGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::IoT::Certificate"
                        },
                        {
                            "$ref": "#/definitions/AWS::IoT::Policy"
                        },
                        {
                            "$ref": "#/definitions/AWS::IoT::PolicyPrincipalAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::IoT::Thing"
                        },
                        {
                            "$ref": "#/definitions/AWS::IoT::ThingPrincipalAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::IoT::TopicRule"
                        },
                        {
                            "$ref": "#/definitions/AWS::KMS::Alias"
                        },
                        {
                            "$ref": "#/definitions/AWS::KMS::Key"
                        },
                        {
                            "$ref": "#/definitions/AWS::Kinesis::Stream"
                        },
                        {
                            "$ref": "#/definitions/AWS::KinesisAnalytics::Application"
                        },
                        {
                            "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationOutput"
                        },
                        {
                            "$ref": "#/definitions/AWS::KinesisAnalytics::ApplicationReferenceDataSource"
                        },
                        {
                            "$ref": "#/definitions/AWS::KinesisFirehose::DeliveryStream"
                        },
                        {
                            "$ref": "#/definitions/AWS::Lambda::Alias"
                        },
                        {
                            "$ref": "#/definitions/AWS::Lambda::EventSourceMapping"
                        },
                        {
                            "$ref": "#/definitions/AWS::Lambda::Function"
                        },
                        {
                            "$ref": "#/definitions/AWS::Lambda::Permission"
                        },
                        {
                            "$ref": "#/definitions/AWS::Lambda::Version"
                        },
                        {
                            "$ref": "#/definitions/AWS::Logs::Destination"
                        },
                        {
                            "$ref": "#/definitions/AWS::Logs::LogGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Logs::LogStream"
                        },
                        {
                            "$ref": "#/definitions/AWS::Logs::MetricFilter"
                        },
                        {
                            "$ref": "#/definitions/AWS::Logs::SubscriptionFilter"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::App"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::ElasticLoadBalancerAttachment"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::Instance"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::Layer"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::Stack"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::UserProfile"
                        },
                        {
                            "$ref": "#/definitions/AWS::OpsWorks::Volume"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBCluster"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBClusterParameterGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBInstance"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBParameterGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBSecurityGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBSecurityGroupIngress"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::DBSubnetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::EventSubscription"
                        },
                        {
                            "$ref": "#/definitions/AWS::RDS::OptionGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Redshift::Cluster"
                        },
                        {
                            "$ref": "#/definitions/AWS::Redshift::ClusterParameterGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Redshift::ClusterSecurityGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Redshift::ClusterSecurityGroupIngress"
                        },
                        {
                            "$ref": "#/definitions/AWS::Redshift::ClusterSubnetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::Route53::HealthCheck"
                        },
                        {
                            "$ref": "#/definitions/AWS::Route53::HostedZone"
                        },
                        {
                            "$ref": "#/definitions/AWS::Route53::RecordSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::Route53::RecordSetGroup"
                        },
                        {
                            "$ref": "#/definitions/AWS::S3::Bucket"
                        },
                        {
                            "$ref": "#/definitions/AWS::S3::BucketPolicy"
                        },
                        {
                            "$ref": "#/definitions/AWS::SDB::Domain"
                        },
                        {
                            "$ref": "#/definitions/AWS::SES::ConfigurationSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::SES::ConfigurationSetEventDestination"
                        },
                        {
                            "$ref": "#/definitions/AWS::SES::ReceiptFilter"
                        },
                        {
                            "$ref": "#/definitions/AWS::SES::ReceiptRule"
                        },
                        {
                            "$ref": "#/definitions/AWS::SES::ReceiptRuleSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::SES::Template"
                        },
                        {
                            "$ref": "#/definitions/AWS::SNS::Subscription"
                        },
                        {
                            "$ref": "#/definitions/AWS::SNS::Topic"
                        },
                        {
                            "$ref": "#/definitions/AWS::SNS::TopicPolicy"
                        },
                        {
                            "$ref": "#/definitions/AWS::SQS::Queue"
                        },
                        {
                            "$ref": "#/definitions/AWS::SQS::QueuePolicy"
                        },
                        {
                            "$ref": "#/definitions/AWS::SSM::Association"
                        },
                        {
                            "$ref": "#/definitions/AWS::SSM::Document"
                        },
                        {
                            "$ref": "#/definitions/AWS::SSM::MaintenanceWindowTask"
                        },
                        {
                            "$ref": "#/definitions/AWS::SSM::Parameter"
                        },
                        {
                            "$ref": "#/definitions/AWS::SSM::PatchBaseline"
                        },
                        {
                            "$ref": "#/definitions/AWS::ServiceDiscovery::Instance"
                        },
                        {
                            "$ref": "#/definitions/AWS::ServiceDiscovery::PrivateDnsNamespace"
                        },
                        {
                            "$ref": "#/definitions/AWS::ServiceDiscovery::PublicDnsNamespace"
                        },
                        {
                            "$ref": "#/definitions/AWS::ServiceDiscovery::Service"
                        },
                        {
                            "$ref": "#/definitions/AWS::StepFunctions::Activity"
                        },
                        {
                            "$ref": "#/definitions/AWS::StepFunctions::StateMachine"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::ByteMatchSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::IPSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::Rule"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::SizeConstraintSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::SqlInjectionMatchSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::WebACL"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAF::XssMatchSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::ByteMatchSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::IPSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::Rule"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::SizeConstraintSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::SqlInjectionMatchSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::WebACL"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::WebACLAssociation"
                        },
                        {
                            "$ref": "#/definitions/AWS::WAFRegional::XssMatchSet"
                        },
                        {
                            "$ref": "#/definitions/AWS::WorkSpaces::Workspace"
                        }
                    ]
                }
            },
            "type": "object"
        },
        "Transform": {
            "type": [
                "object",
                "string"
            ]
        }
    },
    "required": [
        "Resources"
    ],
    "type": "object"
}`
