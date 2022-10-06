// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

// MessageType is the header to bootstrap you way unmarshalling other messages
/*
	Example:
	{
		"type": "message",
		"message": {
			"type": "started_listening"
		}
	}
*/
type MessageType struct {
	Type string `json:"type"`
}

// SybmlMessageType is the header to bootstrap you way unmarshalling other messages
/*
	Example:
	{
		"type": "message",
		"message": {
			"type": "started_listening"
		}
	}
*/
type SybmlMessageType struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
	} `json:"message"`
}

// SymblInitializationMessage the init message when mt.Type == "conversation_created"
/*
	Example:
	{
		"type": "message",
		"message": {
			"type": "conversation_created",
			"data": {
				"conversationId": "5751229838262272"
			}
		}
	}
*/
type SymblInitializationMessage struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
		Data struct {
			ConversationID string `json:"conversationId"`
		} `json:"data"`
	} `json:"message"`
}

// SymblError when mt.Type == "error"
type SymblError struct {
	Type    string `json:"type"`
	Details string `json:"details"`
	Message string `json:"message"`
}

/*
	Conversation Insights
*/

// Testing 1, 2, 3
/*
{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "Tell",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "Tell"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 4683
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "Has",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "Has"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 4724
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "Test",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "Test"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 4730
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "Testing",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "Testing"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 4846
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "Testing 1, 2 3",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "Testing 1, 2 3"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 6065
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": true,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [
              {
                "word": "Testing",
                "startTime": {
                  "seconds": "0",
                  "nanos": "900000000"
                },
                "endTime": {
                  "seconds": "1",
                  "nanos": "600000000"
                }
              },
              {
                "word": "1,",
                "startTime": {
                  "seconds": "1",
                  "nanos": "600000000"
                },
                "endTime": {
                  "seconds": "1",
                  "nanos": "800000000"
                }
              },
              {
                "word": "2",
                "startTime": {
                  "seconds": "1",
                  "nanos": "800000000"
                },
                "endTime": {
                  "seconds": "2",
                  "nanos": "000000000"
                }
              },
              {
                "word": "3.",
                "startTime": {
                  "seconds": "2",
                  "nanos": "000000000"
                },
                "endTime": {
                  "seconds": "2",
                  "nanos": "100000000"
                }
              }
            ],
            "transcript": "Testing 1, 2 3.",
            "confidence": 0.8634329438209534
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "Testing 1, 2 3."
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 6781
}

{
  "type": "message_response",
  "messages": [
    {
      "from": {
        "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe",
        "name": "Jane Doe",
        "userId": "user@email.com"
      },
      "payload": {
        "content": "Testing 1, 2 3.",
        "contentType": "text/plain"
      },
      "id": "a1c3f182-b306-4baf-bf48-ea9659b20097",
      "channel": {
        "id": "realtime-api"
      },
      "metadata": {
        "disablePunctuation": true,
        "timezoneOffset": 0,
        "originalContent": "Testing 1, 2 3.",
        "words": "[{\"word\":\"Testing\",\"startTime\":\"2022-10-05T19:46:52.563Z\",\"endTime\":\"2022-10-05T19:46:53.263Z\",\"timeOffset\":0.9,\"duration\":0.7},{\"word\":\"1,\",\"startTime\":\"2022-10-05T19:46:53.263Z\",\"endTime\":\"2022-10-05T19:46:53.463Z\",\"timeOffset\":1.6,\"duration\":0.2},{\"word\":\"2\",\"startTime\":\"2022-10-05T19:46:53.463Z\",\"endTime\":\"2022-10-05T19:46:53.663Z\",\"timeOffset\":1.8,\"duration\":0.2},{\"word\":\"3.\",\"startTime\":\"2022-10-05T19:46:53.663Z\",\"endTime\":\"2022-10-05T19:46:53.763Z\",\"timeOffset\":2,\"duration\":0.1}]",
        "originalMessageId": "a1c3f182-b306-4baf-bf48-ea9659b20097"
      },
      "dismissed": false,
      "duration": {
        "startTime": "2022-10-05T19:46:52.563Z",
        "endTime": "2022-10-05T19:46:53.763Z",
        "timeOffset": 0.9,
        "duration": 1.2
      }
    }
  ],
  "sequenceNumber": 0
}

// How are you doing today?
{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "How",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "How"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 8628
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "How are",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "How are"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 8639
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "How are you doing today",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "How are you doing today"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 9844
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": true,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [
              {
                "word": "How",
                "startTime": {
                  "seconds": "5",
                  "nanos": "000000000"
                },
                "endTime": {
                  "seconds": "5",
                  "nanos": "300000000"
                }
              },
              {
                "word": "are",
                "startTime": {
                  "seconds": "5",
                  "nanos": "300000000"
                },
                "endTime": {
                  "seconds": "5",
                  "nanos": "400000000"
                }
              },
              {
                "word": "you",
                "startTime": {
                  "seconds": "5",
                  "nanos": "400000000"
                },
                "endTime": {
                  "seconds": "5",
                  "nanos": "600000000"
                }
              },
              {
                "word": "doing",
                "startTime": {
                  "seconds": "5",
                  "nanos": "600000000"
                },
                "endTime": {
                  "seconds": "5",
                  "nanos": "800000000"
                }
              },
              {
                "word": "today?",
                "startTime": {
                  "seconds": "5",
                  "nanos": "800000000"
                },
                "endTime": {
                  "seconds": "6",
                  "nanos": "000000000"
                }
              }
            ],
            "transcript": "How are you doing today?",
            "confidence": 0.9876290559768677
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "How are you doing today?"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 10494
}

{
  "type": "message_response",
  "messages": [
    {
      "from": {
        "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe",
        "name": "Jane Doe",
        "userId": "user@email.com"
      },
      "payload": {
        "content": "How are you doing today?",
        "contentType": "text/plain"
      },
      "id": "63a52ae6-79db-4e4f-97bf-ebea7f1abe6f",
      "channel": {
        "id": "realtime-api"
      },
      "metadata": {
        "disablePunctuation": true,
        "timezoneOffset": 0,
        "originalContent": "How are you doing today?",
        "words": "[{\"word\":\"How\",\"startTime\":\"2022-10-05T19:46:56.663Z\",\"endTime\":\"2022-10-05T19:46:56.963Z\",\"timeOffset\":5,\"duration\":0.3},{\"word\":\"are\",\"startTime\":\"2022-10-05T19:46:56.963Z\",\"endTime\":\"2022-10-05T19:46:57.063Z\",\"timeOffset\":5.3,\"duration\":0.1},{\"word\":\"you\",\"startTime\":\"2022-10-05T19:46:57.063Z\",\"endTime\":\"2022-10-05T19:46:57.263Z\",\"timeOffset\":5.4,\"duration\":0.2},{\"word\":\"doing\",\"startTime\":\"2022-10-05T19:46:57.263Z\",\"endTime\":\"2022-10-05T19:46:57.463Z\",\"timeOffset\":5.6,\"duration\":0.2},{\"word\":\"today?\",\"startTime\":\"2022-10-05T19:46:57.463Z\",\"endTime\":\"2022-10-05T19:46:57.663Z\",\"timeOffset\":5.8,\"duration\":0.2}]",
        "originalMessageId": "63a52ae6-79db-4e4f-97bf-ebea7f1abe6f"
      },
      "dismissed": false,
      "duration": {
        "startTime": "2022-10-05T19:46:56.663Z",
        "endTime": "2022-10-05T19:46:57.663Z",
        "timeOffset": 5,
        "duration": 1
      }
    }
  ],
  "sequenceNumber": 1
}

// insight response... identified question
{
  "type": "insight_response",
  "insights": [
    {
      "id": "9c050e5f-c50a-413c-a44b-6fbc857a62a1",
      "confidence": 0.9794721146232228,
      "hints": [
        {
          "key": "confidenceScore",
          "value": "0.9995726624739847"
        },
        {
          "key": "comprehensionScore",
          "value": "0.9593715667724609"
        }
      ],
      "type": "question",
      "assignee": {
        "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe",
        "name": "Jane Doe",
        "userId": "user@email.com"
      },
      "tags": [],
      "dismissed": false,
      "payload": {
        "content": "How are you doing today?",
        "contentType": "text/plain"
      },
      "from": {
        "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe",
        "name": "Jane Doe",
        "userId": "user@email.com"
      },
      "entities": null,
      "messageReference": {
        "id": "63a52ae6-79db-4e4f-97bf-ebea7f1abe6f"
      }
    }
  ],
  "sequenceNumber": 1
}

// I'm fine
{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "I'm",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "I'm"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 13432
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": false,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [],
            "transcript": "I'm fine",
            "confidence": 0
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "I'm fine"
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 13685
}

{
  "type": "message",
  "message": {
    "type": "recognition_result",
    "isFinal": true,
    "payload": {
      "raw": {
        "alternatives": [
          {
            "words": [
              {
                "word": "I'm",
                "startTime": {
                  "seconds": "9",
                  "nanos": "800000000"
                },
                "endTime": {
                  "seconds": "10",
                  "nanos": "200000000"
                }
              },
              {
                "word": "fine.",
                "startTime": {
                  "seconds": "10",
                  "nanos": "200000000"
                },
                "endTime": {
                  "seconds": "10",
                  "nanos": "500000000"
                }
              },
              {
                "word": "Thanks.",
                "startTime": {
                  "seconds": "10",
                  "nanos": "500000000"
                },
                "endTime": {
                  "seconds": "11",
                  "nanos": "100000000"
                }
              }
            ],
            "transcript": "I'm fine. Thanks.",
            "confidence": 0.9618560671806335
          }
        ]
      }
    },
    "punctuated": {
      "transcript": "I'm fine. Thanks."
    },
    "user": {
      "userId": "user@email.com",
      "name": "Jane Doe",
      "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe"
    }
  },
  "timeOffset": 15494
}

{
  "type": "message_response",
  "messages": [
    {
      "from": {
        "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe",
        "name": "Jane Doe",
        "userId": "user@email.com"
      },
      "payload": {
        "content": "I am fine.",
        "contentType": "text/plain"
      },
      "id": "6d2c7a2e-58c4-4137-b10d-e282c044d945",
      "channel": {
        "id": "realtime-api"
      },
      "metadata": {
        "disablePunctuation": true,
        "timezoneOffset": 0,
        "originalContent": "I'm fine.",
        "words": "[{\"word\":\"I'm\",\"startTime\":\"2022-10-05T19:47:01.463Z\",\"endTime\":\"2022-10-05T19:47:01.863Z\",\"timeOffset\":9.8,\"duration\":0.4},{\"word\":\"fine.\",\"startTime\":\"2022-10-05T19:47:01.863Z\",\"endTime\":\"2022-10-05T19:47:02.163Z\",\"timeOffset\":10.2,\"duration\":0.3}]",
        "originalMessageId": "6d2c7a2e-58c4-4137-b10d-e282c044d945"
      },
      "dismissed": false,
      "duration": {
        "startTime": "2022-10-05T19:47:01.463Z",
        "endTime": "2022-10-05T19:47:02.163Z",
        "timeOffset": 9.8,
        "duration": 0.7
      }
    },
    {
      "from": {
        "id": "bc58c2d6-e4d1-4d9f-a668-86587efc3dfe",
        "name": "Jane Doe",
        "userId": "user@email.com"
      },
      "payload": {
        "content": "Thanks.",
        "contentType": "text/plain"
      },
      "id": "b70d809e-e6d7-4c36-9b1f-f3fd39c13fb6",
      "channel": {
        "id": "realtime-api"
      },
      "metadata": {
        "disablePunctuation": true,
        "timezoneOffset": 0,
        "originalContent": "Thanks.",
        "words": "[{\"word\":\"Thanks.\",\"startTime\":\"2022-10-05T19:47:02.163Z\",\"endTime\":\"2022-10-05T19:47:02.763Z\",\"timeOffset\":10.5,\"duration\":0.6}]",
        "originalMessageId": "b70d809e-e6d7-4c36-9b1f-f3fd39c13fb6"
      },
      "dismissed": false,
      "duration": {
        "startTime": "2022-10-05T19:47:02.163Z",
        "endTime": "2022-10-05T19:47:02.763Z",
        "timeOffset": 10.5,
        "duration": 0.6
      }
    }
  ],
  "sequenceNumber": 2
}


*/
