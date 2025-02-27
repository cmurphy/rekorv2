// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2025 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Rekor is a cryptographically secure, immutable transparency log for signed software releases.",
    "title": "Rekor",
    "version": "1.0.0"
  },
  "host": "rekor.sigstore.dev",
  "paths": {
    "/api/v1/log": {
      "get": {
        "description": "Returns the current root hash and size of the merkle tree used to store the log entries.",
        "tags": [
          "tlog"
        ],
        "summary": "Get information about the current state of the transparency log",
        "operationId": "getLogInfo",
        "parameters": [
          {
            "type": "boolean",
            "default": false,
            "description": "Whether to return a stable checkpoint for the active shard",
            "name": "stable",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A JSON object with the root hash and tree size as properties",
            "schema": {
              "$ref": "#/definitions/LogInfo"
            }
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/api/v1/log/entries": {
      "post": {
        "description": "Creates an entry in the transparency log for a detached signature, public key, and content. Items can be included in the request or fetched by the server when URLs are specified.\n",
        "tags": [
          "entries"
        ],
        "summary": "Creates an entry in the transparency log",
        "operationId": "createLogEntry",
        "parameters": [
          {
            "name": "proposedEntry",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProposedEntry"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Returns the entry created in the transparency log",
            "schema": {
              "$ref": "#/definitions/LogEntry"
            },
            "headers": {
              "ETag": {
                "type": "string",
                "description": "UUID of log entry"
              },
              "Location": {
                "type": "string",
                "format": "uri",
                "description": "URI location of log entry"
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadContent"
          },
          "409": {
            "$ref": "#/responses/Conflict"
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/api/v1/log/publicKey": {
      "get": {
        "description": "Returns the public key that can be used to validate the signed tree head",
        "produces": [
          "application/x-pem-file"
        ],
        "tags": [
          "pubkey"
        ],
        "summary": "Retrieve the public key that can be used to validate the signed tree head",
        "operationId": "getPublicKey",
        "parameters": [
          {
            "pattern": "^[0-9]+$",
            "type": "string",
            "description": "The tree ID of the tree you wish to get a public key for",
            "name": "treeID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The public key",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "InactiveShardLogInfo": {
      "type": "object",
      "required": [
        "rootHash",
        "treeSize",
        "signedTreeHead",
        "treeID"
      ],
      "properties": {
        "rootHash": {
          "description": "The current hash value stored at the root of the merkle tree",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "signedTreeHead": {
          "description": "The current signed tree head",
          "type": "string",
          "format": "signedCheckpoint"
        },
        "treeID": {
          "description": "The current treeID",
          "type": "string",
          "pattern": "^[0-9]+$"
        },
        "treeSize": {
          "description": "The current number of nodes in the merkle tree",
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "InclusionProof": {
      "type": "object",
      "required": [
        "logIndex",
        "rootHash",
        "treeSize",
        "hashes",
        "checkpoint"
      ],
      "properties": {
        "checkpoint": {
          "description": "The checkpoint (signed tree head) that the inclusion proof is based on",
          "type": "string",
          "format": "signedCheckpoint"
        },
        "hashes": {
          "description": "A list of hashes required to compute the inclusion proof, sorted in order from leaf to root",
          "type": "array",
          "items": {
            "description": "SHA256 hash value expressed in hexadecimal format",
            "type": "string",
            "pattern": "^[0-9a-fA-F]{64}$"
          }
        },
        "logIndex": {
          "description": "The index of the entry in the transparency log",
          "type": "integer"
        },
        "rootHash": {
          "description": "The hash value stored at the root of the merkle tree at the time the proof was generated",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "treeSize": {
          "description": "The size of the merkle tree at the time the inclusion proof was generated",
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "LogEntry": {
      "type": "object",
      "additionalProperties": {
        "type": "object",
        "required": [
          "logID",
          "logIndex",
          "body",
          "integratedTime"
        ],
        "properties": {
          "attestation": {
            "type": "object",
            "format": "byte",
            "properties": {
              "data": {
                "format": "byte"
              }
            }
          },
          "body": {
            "type": "object",
            "additionalProperties": true
          },
          "integratedTime": {
            "description": "The time the entry was added to the log as a Unix timestamp in seconds",
            "type": "integer"
          },
          "logID": {
            "description": "This is the SHA256 hash of the DER-encoded public key for the log at the time the entry was included in the log",
            "type": "string",
            "pattern": "^[0-9a-fA-F]{64}$"
          },
          "logIndex": {
            "type": "integer"
          },
          "verification": {
            "type": "object",
            "properties": {
              "inclusionProof": {
                "$ref": "#/definitions/InclusionProof"
              },
              "signedEntryTimestamp": {
                "description": "Signature over the logID, logIndex, body and integratedTime.",
                "type": "string",
                "format": "byte"
              }
            }
          }
        }
      }
    },
    "LogInfo": {
      "type": "object",
      "required": [
        "rootHash",
        "treeSize",
        "signedTreeHead",
        "treeID"
      ],
      "properties": {
        "inactiveShards": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/InactiveShardLogInfo"
          }
        },
        "rootHash": {
          "description": "The current hash value stored at the root of the merkle tree",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "signedTreeHead": {
          "description": "The current signed tree head",
          "type": "string",
          "format": "signedCheckpoint"
        },
        "treeID": {
          "description": "The current treeID",
          "type": "string",
          "pattern": "^[0-9]+$"
        },
        "treeSize": {
          "description": "The current number of nodes in the merkle tree",
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "ProposedEntry": {
      "type": "object",
      "required": [
        "kind"
      ],
      "properties": {
        "kind": {
          "type": "string"
        }
      },
      "discriminator": "kind"
    },
    "dsse": {
      "description": "DSSE envelope",
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/ProposedEntry"
        },
        {
          "required": [
            "apiVersion",
            "spec"
          ],
          "properties": {
            "apiVersion": {
              "type": "string",
              "pattern": "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
            },
            "spec": {
              "type": "object",
              "$ref": "pkg/types/dsse/dsse_schema.json"
            }
          },
          "additionalProperties": false
        }
      ]
    },
    "hashedrekord": {
      "description": "Hashed Rekord object",
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/ProposedEntry"
        },
        {
          "required": [
            "apiVersion",
            "spec"
          ],
          "properties": {
            "apiVersion": {
              "type": "string",
              "pattern": "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
            },
            "spec": {
              "type": "object",
              "$ref": "pkg/types/hashedrekord/hashedrekord_schema.json"
            }
          },
          "additionalProperties": false
        }
      ]
    }
  },
  "responses": {
    "BadContent": {
      "description": "The content supplied to the server was invalid",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Conflict": {
      "description": "The request conflicts with the current state of the transparency log",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Location": {
          "type": "string",
          "format": "uri"
        }
      }
    },
    "InternalServerError": {
      "description": "There was an internal error in the server while processing the request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "The content requested could not be found"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Rekor is a cryptographically secure, immutable transparency log for signed software releases.",
    "title": "Rekor",
    "version": "1.0.0"
  },
  "host": "rekor.sigstore.dev",
  "paths": {
    "/api/v1/log": {
      "get": {
        "description": "Returns the current root hash and size of the merkle tree used to store the log entries.",
        "tags": [
          "tlog"
        ],
        "summary": "Get information about the current state of the transparency log",
        "operationId": "getLogInfo",
        "parameters": [
          {
            "type": "boolean",
            "default": false,
            "description": "Whether to return a stable checkpoint for the active shard",
            "name": "stable",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A JSON object with the root hash and tree size as properties",
            "schema": {
              "$ref": "#/definitions/LogInfo"
            }
          },
          "default": {
            "description": "There was an internal error in the server while processing the request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/log/entries": {
      "post": {
        "description": "Creates an entry in the transparency log for a detached signature, public key, and content. Items can be included in the request or fetched by the server when URLs are specified.\n",
        "tags": [
          "entries"
        ],
        "summary": "Creates an entry in the transparency log",
        "operationId": "createLogEntry",
        "parameters": [
          {
            "name": "proposedEntry",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProposedEntry"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Returns the entry created in the transparency log",
            "schema": {
              "$ref": "#/definitions/LogEntry"
            },
            "headers": {
              "ETag": {
                "type": "string",
                "description": "UUID of log entry"
              },
              "Location": {
                "type": "string",
                "format": "uri",
                "description": "URI location of log entry"
              }
            }
          },
          "400": {
            "description": "The content supplied to the server was invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "The request conflicts with the current state of the transparency log",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "Location": {
                "type": "string",
                "format": "uri"
              }
            }
          },
          "default": {
            "description": "There was an internal error in the server while processing the request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/log/publicKey": {
      "get": {
        "description": "Returns the public key that can be used to validate the signed tree head",
        "produces": [
          "application/x-pem-file"
        ],
        "tags": [
          "pubkey"
        ],
        "summary": "Retrieve the public key that can be used to validate the signed tree head",
        "operationId": "getPublicKey",
        "parameters": [
          {
            "pattern": "^[0-9]+$",
            "type": "string",
            "description": "The tree ID of the tree you wish to get a public key for",
            "name": "treeID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The public key",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "There was an internal error in the server while processing the request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DSSEV001SchemaEnvelopeHash": {
      "description": "Specifies the hash algorithm and value encompassing the entire envelope sent to Rekor",
      "type": "object",
      "required": [
        "algorithm",
        "value"
      ],
      "properties": {
        "algorithm": {
          "description": "The hashing function used to compute the hash value",
          "type": "string",
          "enum": [
            "sha256"
          ]
        },
        "value": {
          "description": "The value of the computed digest over the entire envelope",
          "type": "string"
        }
      },
      "readOnly": true
    },
    "DSSEV001SchemaPayloadHash": {
      "description": "Specifies the hash algorithm and value covering the payload within the DSSE envelope",
      "type": "object",
      "required": [
        "algorithm",
        "value"
      ],
      "properties": {
        "algorithm": {
          "description": "The hashing function used to compute the hash value",
          "type": "string",
          "enum": [
            "sha256"
          ]
        },
        "value": {
          "description": "The value of the computed digest over the payload within the envelope",
          "type": "string"
        }
      },
      "readOnly": true
    },
    "DSSEV001SchemaProposedContent": {
      "type": "object",
      "required": [
        "envelope",
        "verifiers"
      ],
      "properties": {
        "envelope": {
          "description": "DSSE envelope specified as a stringified JSON object",
          "type": "string",
          "writeOnly": true
        },
        "verifiers": {
          "description": "collection of all verification material (e.g. public keys or certificates) used to verify signatures over envelope's payload, specified as base64-encoded strings",
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string",
            "format": "byte"
          },
          "writeOnly": true
        }
      },
      "writeOnly": true
    },
    "DSSEV001SchemaSignaturesItems0": {
      "description": "a signature of the envelope's payload along with the verification material for the signature",
      "type": "object",
      "required": [
        "signature",
        "verifier"
      ],
      "properties": {
        "signature": {
          "description": "base64 encoded signature of the payload",
          "type": "string",
          "pattern": "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
        },
        "verifier": {
          "description": "verification material that was used to verify the corresponding signature, specified as a base64 encoded string",
          "type": "string",
          "format": "byte"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "HashedrekordV001SchemaData": {
      "description": "Information about the content associated with the entry",
      "type": "object",
      "properties": {
        "hash": {
          "description": "Specifies the hash algorithm and value for the content",
          "type": "object",
          "required": [
            "algorithm",
            "value"
          ],
          "properties": {
            "algorithm": {
              "description": "The hashing function used to compute the hash value",
              "type": "string",
              "enum": [
                "sha256",
                "sha384",
                "sha512"
              ]
            },
            "value": {
              "description": "The hash value for the content, as represented by a lower case hexadecimal string",
              "type": "string"
            }
          }
        }
      }
    },
    "HashedrekordV001SchemaDataHash": {
      "description": "Specifies the hash algorithm and value for the content",
      "type": "object",
      "required": [
        "algorithm",
        "value"
      ],
      "properties": {
        "algorithm": {
          "description": "The hashing function used to compute the hash value",
          "type": "string",
          "enum": [
            "sha256",
            "sha384",
            "sha512"
          ]
        },
        "value": {
          "description": "The hash value for the content, as represented by a lower case hexadecimal string",
          "type": "string"
        }
      }
    },
    "HashedrekordV001SchemaSignature": {
      "description": "Information about the detached signature associated with the entry",
      "type": "object",
      "properties": {
        "content": {
          "description": "Specifies the content of the signature inline within the document",
          "type": "string",
          "format": "byte"
        },
        "publicKey": {
          "description": "The public key that can verify the signature; this can also be an X509 code signing certificate that contains the raw public key information",
          "type": "object",
          "properties": {
            "content": {
              "description": "Specifies the content of the public key or code signing certificate inline within the document",
              "type": "string",
              "format": "byte"
            }
          }
        }
      }
    },
    "HashedrekordV001SchemaSignaturePublicKey": {
      "description": "The public key that can verify the signature; this can also be an X509 code signing certificate that contains the raw public key information",
      "type": "object",
      "properties": {
        "content": {
          "description": "Specifies the content of the public key or code signing certificate inline within the document",
          "type": "string",
          "format": "byte"
        }
      }
    },
    "InactiveShardLogInfo": {
      "type": "object",
      "required": [
        "rootHash",
        "treeSize",
        "signedTreeHead",
        "treeID"
      ],
      "properties": {
        "rootHash": {
          "description": "The current hash value stored at the root of the merkle tree",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "signedTreeHead": {
          "description": "The current signed tree head",
          "type": "string",
          "format": "signedCheckpoint"
        },
        "treeID": {
          "description": "The current treeID",
          "type": "string",
          "pattern": "^[0-9]+$"
        },
        "treeSize": {
          "description": "The current number of nodes in the merkle tree",
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "InclusionProof": {
      "type": "object",
      "required": [
        "logIndex",
        "rootHash",
        "treeSize",
        "hashes",
        "checkpoint"
      ],
      "properties": {
        "checkpoint": {
          "description": "The checkpoint (signed tree head) that the inclusion proof is based on",
          "type": "string",
          "format": "signedCheckpoint"
        },
        "hashes": {
          "description": "A list of hashes required to compute the inclusion proof, sorted in order from leaf to root",
          "type": "array",
          "items": {
            "description": "SHA256 hash value expressed in hexadecimal format",
            "type": "string",
            "pattern": "^[0-9a-fA-F]{64}$"
          }
        },
        "logIndex": {
          "description": "The index of the entry in the transparency log",
          "type": "integer",
          "minimum": 0
        },
        "rootHash": {
          "description": "The hash value stored at the root of the merkle tree at the time the proof was generated",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "treeSize": {
          "description": "The size of the merkle tree at the time the inclusion proof was generated",
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "LogEntry": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/LogEntryAnon"
      }
    },
    "LogEntryAnon": {
      "type": "object",
      "required": [
        "logID",
        "logIndex",
        "body",
        "integratedTime"
      ],
      "properties": {
        "attestation": {
          "type": "object",
          "format": "byte",
          "properties": {
            "data": {
              "format": "byte"
            }
          }
        },
        "body": {
          "type": "object",
          "additionalProperties": true
        },
        "integratedTime": {
          "description": "The time the entry was added to the log as a Unix timestamp in seconds",
          "type": "integer"
        },
        "logID": {
          "description": "This is the SHA256 hash of the DER-encoded public key for the log at the time the entry was included in the log",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "logIndex": {
          "type": "integer",
          "minimum": 0
        },
        "verification": {
          "type": "object",
          "properties": {
            "inclusionProof": {
              "$ref": "#/definitions/InclusionProof"
            },
            "signedEntryTimestamp": {
              "description": "Signature over the logID, logIndex, body and integratedTime.",
              "type": "string",
              "format": "byte"
            }
          }
        }
      }
    },
    "LogEntryAnonAttestation": {
      "type": "object",
      "format": "byte",
      "properties": {
        "data": {
          "format": "byte"
        }
      }
    },
    "LogEntryAnonVerification": {
      "type": "object",
      "properties": {
        "inclusionProof": {
          "$ref": "#/definitions/InclusionProof"
        },
        "signedEntryTimestamp": {
          "description": "Signature over the logID, logIndex, body and integratedTime.",
          "type": "string",
          "format": "byte"
        }
      }
    },
    "LogInfo": {
      "type": "object",
      "required": [
        "rootHash",
        "treeSize",
        "signedTreeHead",
        "treeID"
      ],
      "properties": {
        "inactiveShards": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/InactiveShardLogInfo"
          }
        },
        "rootHash": {
          "description": "The current hash value stored at the root of the merkle tree",
          "type": "string",
          "pattern": "^[0-9a-fA-F]{64}$"
        },
        "signedTreeHead": {
          "description": "The current signed tree head",
          "type": "string",
          "format": "signedCheckpoint"
        },
        "treeID": {
          "description": "The current treeID",
          "type": "string",
          "pattern": "^[0-9]+$"
        },
        "treeSize": {
          "description": "The current number of nodes in the merkle tree",
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "ProposedEntry": {
      "type": "object",
      "required": [
        "kind"
      ],
      "properties": {
        "kind": {
          "type": "string"
        }
      },
      "discriminator": "kind"
    },
    "dsse": {
      "description": "DSSE envelope",
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/ProposedEntry"
        },
        {
          "required": [
            "apiVersion",
            "spec"
          ],
          "properties": {
            "apiVersion": {
              "type": "string",
              "pattern": "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
            },
            "spec": {
              "$ref": "#/definitions/dsseSchema"
            }
          },
          "additionalProperties": false
        }
      ]
    },
    "dsseSchema": {
      "description": "log entry schema for dsse envelopes",
      "type": "object",
      "title": "DSSE Schema",
      "oneOf": [
        {
          "$ref": "#/definitions/dsseV001Schema"
        }
      ],
      "$schema": "http://json-schema.org/draft-07/schema",
      "$id": "http://rekor.sigstore.dev/types/dsse/dsse_schema.json"
    },
    "dsseV001Schema": {
      "description": "Schema for DSSE envelopes",
      "type": "object",
      "title": "DSSE v0.0.1 Schema",
      "oneOf": [
        {
          "required": [
            "proposedContent"
          ]
        },
        {
          "required": [
            "signatures",
            "envelopeHash",
            "payloadHash"
          ]
        }
      ],
      "properties": {
        "envelopeHash": {
          "description": "Specifies the hash algorithm and value encompassing the entire envelope sent to Rekor",
          "type": "object",
          "required": [
            "algorithm",
            "value"
          ],
          "properties": {
            "algorithm": {
              "description": "The hashing function used to compute the hash value",
              "type": "string",
              "enum": [
                "sha256"
              ]
            },
            "value": {
              "description": "The value of the computed digest over the entire envelope",
              "type": "string"
            }
          },
          "readOnly": true
        },
        "payloadHash": {
          "description": "Specifies the hash algorithm and value covering the payload within the DSSE envelope",
          "type": "object",
          "required": [
            "algorithm",
            "value"
          ],
          "properties": {
            "algorithm": {
              "description": "The hashing function used to compute the hash value",
              "type": "string",
              "enum": [
                "sha256"
              ]
            },
            "value": {
              "description": "The value of the computed digest over the payload within the envelope",
              "type": "string"
            }
          },
          "readOnly": true
        },
        "proposedContent": {
          "type": "object",
          "required": [
            "envelope",
            "verifiers"
          ],
          "properties": {
            "envelope": {
              "description": "DSSE envelope specified as a stringified JSON object",
              "type": "string",
              "writeOnly": true
            },
            "verifiers": {
              "description": "collection of all verification material (e.g. public keys or certificates) used to verify signatures over envelope's payload, specified as base64-encoded strings",
              "type": "array",
              "minItems": 1,
              "items": {
                "type": "string",
                "format": "byte"
              },
              "writeOnly": true
            }
          },
          "writeOnly": true
        },
        "signatures": {
          "description": "extracted collection of all signatures of the envelope's payload; elements will be sorted by lexicographical order of the base64 encoded signature strings",
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/DSSEV001SchemaSignaturesItems0"
          },
          "readOnly": true
        }
      },
      "$schema": "http://json-schema.org/draft-07/schema",
      "$id": "http://rekor.sigstore.dev/types/dsse/dsse_v0_0_1_schema.json"
    },
    "hashedrekord": {
      "description": "Hashed Rekord object",
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/ProposedEntry"
        },
        {
          "required": [
            "apiVersion",
            "spec"
          ],
          "properties": {
            "apiVersion": {
              "type": "string",
              "pattern": "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
            },
            "spec": {
              "$ref": "#/definitions/hashedrekordSchema"
            }
          },
          "additionalProperties": false
        }
      ]
    },
    "hashedrekordSchema": {
      "description": "Schema for Hashedrekord objects",
      "type": "object",
      "title": "Hashedrekord Schema",
      "oneOf": [
        {
          "$ref": "#/definitions/hashedrekordV001Schema"
        }
      ],
      "$schema": "http://json-schema.org/draft-07/schema",
      "$id": "http://rekor.sigstore.dev/types/hashedrekord/hasehedrekord_schema.json"
    },
    "hashedrekordV001Schema": {
      "description": "Schema for Hashed Rekord object",
      "type": "object",
      "title": "Hashed Rekor v0.0.1 Schema",
      "required": [
        "signature",
        "data"
      ],
      "properties": {
        "data": {
          "description": "Information about the content associated with the entry",
          "type": "object",
          "properties": {
            "hash": {
              "description": "Specifies the hash algorithm and value for the content",
              "type": "object",
              "required": [
                "algorithm",
                "value"
              ],
              "properties": {
                "algorithm": {
                  "description": "The hashing function used to compute the hash value",
                  "type": "string",
                  "enum": [
                    "sha256",
                    "sha384",
                    "sha512"
                  ]
                },
                "value": {
                  "description": "The hash value for the content, as represented by a lower case hexadecimal string",
                  "type": "string"
                }
              }
            }
          }
        },
        "signature": {
          "description": "Information about the detached signature associated with the entry",
          "type": "object",
          "properties": {
            "content": {
              "description": "Specifies the content of the signature inline within the document",
              "type": "string",
              "format": "byte"
            },
            "publicKey": {
              "description": "The public key that can verify the signature; this can also be an X509 code signing certificate that contains the raw public key information",
              "type": "object",
              "properties": {
                "content": {
                  "description": "Specifies the content of the public key or code signing certificate inline within the document",
                  "type": "string",
                  "format": "byte"
                }
              }
            }
          }
        }
      },
      "$schema": "http://json-schema.org/draft-07/schema",
      "$id": "http://rekor.sigstore.dev/types/rekord/hashedrekord_v0_0_1_schema.json"
    }
  },
  "responses": {
    "BadContent": {
      "description": "The content supplied to the server was invalid",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Conflict": {
      "description": "The request conflicts with the current state of the transparency log",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Location": {
          "type": "string",
          "format": "uri"
        }
      }
    },
    "InternalServerError": {
      "description": "There was an internal error in the server while processing the request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "The content requested could not be found"
    }
  }
}`))
}
