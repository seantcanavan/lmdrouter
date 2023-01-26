# Lambda JWT Router
Simple HTTP router for working with Json Web Tokens (JWTs) on AWS Lambda through APIGateway proxy requests.
Supports local HTTP routing through net/http for local testing, debugging, and development.
Supports setting an HMAC secret for signing and verifying JWTs.
Supports automatically replying to HTTP OPTIONS requests so calls from browsers succeed.

## Previous README
Go HTTP router library for AWS API Gateway-invoked Lambda Functions
Forked from [aquasecurity/lmdrouter](https://github.com/aquasecurity/lmdrouter)


## How to Build locally
1. `make build`

## How to Test locally
1. `make test`

## How to Use
1. TBD

## Sample routing example

## Sample JWT example

## Sample middleware example

## All tests are passing
```
=== RUN   TestAllowOptionsMW
=== RUN   TestAllowOptionsMW/verify_empty_OPTIONS_req_succeeds
=== RUN   TestAllowOptionsMW/verify_OPTIONS_req_succeeds_with_invalid_JWT_for_AllowOptions
=== RUN   TestAllowOptionsMW/verify_OPTIONS_req_succeeds_with_no_Authorization_header_for_AllowOptions
--- PASS: TestAllowOptionsMW (0.00s)
    --- PASS: TestAllowOptionsMW/verify_empty_OPTIONS_req_succeeds (0.00s)
    --- PASS: TestAllowOptionsMW/verify_OPTIONS_req_succeeds_with_invalid_JWT_for_AllowOptions (0.00s)
    --- PASS: TestAllowOptionsMW/verify_OPTIONS_req_succeeds_with_no_Authorization_header_for_AllowOptions (0.00s)
=== RUN   TestDecodeAndInjectExpandedClaims
=== RUN   TestDecodeAndInjectExpandedClaims/verify_error_is_returned_by_DecodeExpanded_when_missing_Authorization_header
=== RUN   TestDecodeAndInjectExpandedClaims/verify_context_is_returned_by_DecodeExpanded_with_a_signed_JWT
--- PASS: TestDecodeAndInjectExpandedClaims (0.00s)
    --- PASS: TestDecodeAndInjectExpandedClaims/verify_error_is_returned_by_DecodeExpanded_when_missing_Authorization_header (0.00s)
    --- PASS: TestDecodeAndInjectExpandedClaims/verify_context_is_returned_by_DecodeExpanded_with_a_signed_JWT (0.00s)
=== RUN   TestDecodeAndInjectStandardClaims
=== RUN   TestDecodeAndInjectStandardClaims/verify_error_is_returned_by_DecodeStandard_when_missing_Authorization_header
=== RUN   TestDecodeAndInjectStandardClaims/verify_context_is_returned_by_DecodeStandard_with_a_signed_JWT
--- PASS: TestDecodeAndInjectStandardClaims (0.00s)
    --- PASS: TestDecodeAndInjectStandardClaims/verify_error_is_returned_by_DecodeStandard_when_missing_Authorization_header (0.00s)
    --- PASS: TestDecodeAndInjectStandardClaims/verify_context_is_returned_by_DecodeStandard_with_a_signed_JWT (0.00s)
=== RUN   TestExtractJWT
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_empty_Authorization_header
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_Authorization_header_misspelled_-_all_caps
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_Authorization_header_misspelled_-_lowercase
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_bearer_prefix_not_used
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_bearer_not_camel_cased
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_BEARER_all_caps
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_err_for_Bearer_does_not_end_with_space
=== RUN   TestExtractJWT/verify_ExtractJWT_returns_claims_correctly_with_valid_input
--- PASS: TestExtractJWT (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_empty_Authorization_header (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_Authorization_header_misspelled_-_all_caps (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_Authorization_header_misspelled_-_lowercase (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_bearer_prefix_not_used (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_bearer_not_camel_cased (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_BEARER_all_caps (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_err_for_Bearer_does_not_end_with_space (0.00s)
    --- PASS: TestExtractJWT/verify_ExtractJWT_returns_claims_correctly_with_valid_input (0.00s)
=== RUN   TestGenerateEmptyErrorHandler
=== RUN   TestGenerateEmptyErrorHandler/verify_empty_error_handler_returns_error
--- PASS: TestGenerateEmptyErrorHandler (0.00s)
    --- PASS: TestGenerateEmptyErrorHandler/verify_empty_error_handler_returns_error (0.00s)
=== RUN   TestGenerateEmptySuccessHandler
=== RUN   TestGenerateEmptySuccessHandler/verify_empty_success_handler_returns_success
--- PASS: TestGenerateEmptySuccessHandler (0.00s)
    --- PASS: TestGenerateEmptySuccessHandler/verify_empty_success_handler_returns_success (0.00s)
=== RUN   TestExtendExpandedClaims
=== RUN   TestExtendExpandedClaims/verify_sign_and_verify_expanded_and_custom_fields_in_claims
--- PASS: TestExtendExpandedClaims (0.00s)
    --- PASS: TestExtendExpandedClaims/verify_sign_and_verify_expanded_and_custom_fields_in_claims (0.00s)
=== RUN   TestExtendStandardClaims
=== RUN   TestExtendStandardClaims/verify_sign_and_verify_standard_and_custom_fields_in_claims
--- PASS: TestExtendStandardClaims (0.00s)
    --- PASS: TestExtendStandardClaims/verify_sign_and_verify_standard_and_custom_fields_in_claims (0.00s)
=== RUN   TestExtractCustomClaims
=== RUN   TestExtractCustomClaims/verify_ExtractCustom_returns_an_err_when_unmarshalling_to_invalid_custom_claims_object
=== RUN   TestExtractCustomClaims/verify_ExtractCustom_works_when_called_with_the_correct_parameters
--- PASS: TestExtractCustomClaims (0.00s)
    --- PASS: TestExtractCustomClaims/verify_ExtractCustom_returns_an_err_when_unmarshalling_to_invalid_custom_claims_object (0.00s)
    --- PASS: TestExtractCustomClaims/verify_ExtractCustom_works_when_called_with_the_correct_parameters (0.00s)
=== RUN   TestExtractStandardClaims
=== RUN   TestExtractStandardClaims/verify_ExtractStandard_returns_an_err_when_unmarshalling_to_invalid_standard_claims_object
=== RUN   TestExtractStandardClaims/verify_ExtractCustom_works_when_called_with_the_correct_parameters
--- PASS: TestExtractStandardClaims (0.00s)
    --- PASS: TestExtractStandardClaims/verify_ExtractStandard_returns_an_err_when_unmarshalling_to_invalid_standard_claims_object (0.00s)
    --- PASS: TestExtractStandardClaims/verify_ExtractCustom_works_when_called_with_the_correct_parameters (0.00s)
=== RUN   TestSign
=== RUN   TestSign/verify_signed_jwt_secret_with_valid_standard_claim
--- PASS: TestSign (0.00s)
    --- PASS: TestSign/verify_signed_jwt_secret_with_valid_standard_claim (0.00s)
=== RUN   TestVerifyJWT
=== RUN   TestVerifyJWT/verify_err_when_parsing_invalid_jwt
=== RUN   TestVerifyJWT/verify_err_when_parsing_expired_token_with_valid_jwt
--- PASS: TestVerifyJWT (0.00s)
    --- PASS: TestVerifyJWT/verify_err_when_parsing_invalid_jwt (0.00s)
    --- PASS: TestVerifyJWT/verify_err_when_parsing_expired_token_with_valid_jwt (0.00s)
PASS
ok  	github.com/seantcanavan/lambda_jwt_router/lambda_jwt	0.004s
=== RUN   TestMarshalLambdaRequest
=== RUN   TestMarshalLambdaRequest/verify_MarshalReq_correctly_adds_the_JSON_string_to_the_request_body
--- PASS: TestMarshalLambdaRequest (0.00s)
    --- PASS: TestMarshalLambdaRequest/verify_MarshalReq_correctly_adds_the_JSON_string_to_the_request_body (0.00s)
=== RUN   Test_UnmarshalReq
=== RUN   Test_UnmarshalReq/valid_path&query_input
=== RUN   Test_UnmarshalReq/invalid_path&query_input
=== RUN   Test_UnmarshalReq/valid_body_input,_not_base64
=== RUN   Test_UnmarshalReq/invalid_body_input,_not_base64
=== RUN   Test_UnmarshalReq/valid_body_input,_base64
=== RUN   Test_UnmarshalReq/invalid_body_input,_base64
--- PASS: Test_UnmarshalReq (0.00s)
    --- PASS: Test_UnmarshalReq/valid_path&query_input (0.00s)
    --- PASS: Test_UnmarshalReq/invalid_path&query_input (0.00s)
    --- PASS: Test_UnmarshalReq/valid_body_input,_not_base64 (0.00s)
    --- PASS: Test_UnmarshalReq/invalid_body_input,_not_base64 (0.00s)
    --- PASS: Test_UnmarshalReq/valid_body_input,_base64 (0.00s)
    --- PASS: Test_UnmarshalReq/invalid_body_input,_base64 (0.00s)
=== RUN   TestHTTPHandler
=== RUN   TestHTTPHandler/POST_/api_without_auth
=== RUN   TestHTTPHandler/POST_/api_with_auth
=== RUN   TestHTTPHandler/GET_/api
=== RUN   TestHTTPHandler/GET_/api/something/stuff
--- PASS: TestHTTPHandler (0.00s)
    --- PASS: TestHTTPHandler/POST_/api_without_auth (0.00s)
    --- PASS: TestHTTPHandler/POST_/api_with_auth (0.00s)
    --- PASS: TestHTTPHandler/GET_/api (0.00s)
    --- PASS: TestHTTPHandler/GET_/api/something/stuff (0.00s)
=== RUN   TestCustom
=== RUN   TestCustom/verify_CustomRes_returns_the_struct_in_the_response_body
=== RUN   TestCustom/verify_CustomRes_returns_the_key_value_pair_in_the_response_headers
=== RUN   TestCustom/verify_CustomRes_returns_the_correct_status_code
=== RUN   TestCustom/verify_CustomRes_embeds_CORS_headers_in_the_response_headers
--- PASS: TestCustom (0.00s)
    --- PASS: TestCustom/verify_CustomRes_returns_the_struct_in_the_response_body (0.00s)
    --- PASS: TestCustom/verify_CustomRes_returns_the_key_value_pair_in_the_response_headers (0.00s)
    --- PASS: TestCustom/verify_CustomRes_returns_the_correct_status_code (0.00s)
    --- PASS: TestCustom/verify_CustomRes_embeds_CORS_headers_in_the_response_headers (0.00s)
=== RUN   TestEmpty
=== RUN   TestEmpty/verify_EmptyRes_returns_the_correct_status_code
=== RUN   TestEmpty/verify_EmptyRes_embeds_CORS_headers_in_the_response_headers
--- PASS: TestEmpty (0.00s)
    --- PASS: TestEmpty/verify_EmptyRes_returns_the_correct_status_code (0.00s)
    --- PASS: TestEmpty/verify_EmptyRes_embeds_CORS_headers_in_the_response_headers (0.00s)
=== RUN   TestError
=== RUN   TestError/Handle_an_ErrorAndStatusRes
=== RUN   TestError/Handle_an_ErrorAndStatusRes_when_ExposeServerErrors_is_true
=== RUN   TestError/Handle_an_ErrorAndStatusRes_when_ExposeServerErrors_is_false
=== RUN   TestError/Handle_a_general_error_when_ExposeServerErrors_is_true
=== RUN   TestError/Handle_a_general_error_when_ExposeServerErrors_is_false
--- PASS: TestError (0.00s)
    --- PASS: TestError/Handle_an_ErrorAndStatusRes (0.00s)
    --- PASS: TestError/Handle_an_ErrorAndStatusRes_when_ExposeServerErrors_is_true (0.00s)
    --- PASS: TestError/Handle_an_ErrorAndStatusRes_when_ExposeServerErrors_is_false (0.00s)
    --- PASS: TestError/Handle_a_general_error_when_ExposeServerErrors_is_true (0.00s)
    --- PASS: TestError/Handle_a_general_error_when_ExposeServerErrors_is_false (0.00s)
=== RUN   TestErrorAndStatus
=== RUN   TestErrorAndStatus/verify_ErrorAndStatusRes_returns_the_correct_status_code
=== RUN   TestErrorAndStatus/verify_ErrorAndStatusRes_embeds_CORS_headers_in_the_response
--- PASS: TestErrorAndStatus (0.00s)
    --- PASS: TestErrorAndStatus/verify_ErrorAndStatusRes_returns_the_correct_status_code (0.00s)
    --- PASS: TestErrorAndStatus/verify_ErrorAndStatusRes_embeds_CORS_headers_in_the_response (0.00s)
=== RUN   TestFile
=== RUN   TestFile/verify_FileRes_returns_the_correct_status_code
=== RUN   TestFile/verify_FileRes_marks_the_response_as_NOT_base64_encoded
=== RUN   TestFile/verify_FileRes_embeds_the_bytes_correctly_in_the_response_object_as_a_string
=== RUN   TestFile/verify_FileRes_preserves_the_original_header_values
--- PASS: TestFile (0.00s)
    --- PASS: TestFile/verify_FileRes_returns_the_correct_status_code (0.00s)
    --- PASS: TestFile/verify_FileRes_marks_the_response_as_NOT_base64_encoded (0.00s)
    --- PASS: TestFile/verify_FileRes_embeds_the_bytes_correctly_in_the_response_object_as_a_string (0.00s)
    --- PASS: TestFile/verify_FileRes_preserves_the_original_header_values (0.00s)
=== RUN   TestFileB64
=== RUN   TestFileB64/verify_FileB64Res_returns_the_correct_status_code
=== RUN   TestFileB64/verify_FileB64Res_marks_the_response_as_base64_encoded
=== RUN   TestFileB64/verify_FileB64Res_embeds_the_bytes_correctly_in_the_response_object_as_a_byte64_encoded_string
=== RUN   TestFileB64/verify_FileRes_preserves_the_original_header_values
--- PASS: TestFileB64 (0.00s)
    --- PASS: TestFileB64/verify_FileB64Res_returns_the_correct_status_code (0.00s)
    --- PASS: TestFileB64/verify_FileB64Res_marks_the_response_as_base64_encoded (0.00s)
    --- PASS: TestFileB64/verify_FileB64Res_embeds_the_bytes_correctly_in_the_response_object_as_a_byte64_encoded_string (0.00s)
    --- PASS: TestFileB64/verify_FileRes_preserves_the_original_header_values (0.00s)
=== RUN   TestSuccess
=== RUN   TestSuccess/verify_SuccessRes_returns_the_correct_status_code
=== RUN   TestSuccess/verify_SuccessRes_returns_the_struct_in_the_response_body
--- PASS: TestSuccess (0.00s)
    --- PASS: TestSuccess/verify_SuccessRes_returns_the_correct_status_code (0.00s)
    --- PASS: TestSuccess/verify_SuccessRes_returns_the_struct_in_the_response_body (0.00s)
=== RUN   TestRouter
=== RUN   TestRouter/Routes_created_correctly
=== RUN   TestRouter/Routes_created_correctly//
=== RUN   TestRouter/Routes_created_correctly//:id
=== RUN   TestRouter/Routes_created_correctly//:id/stuff/:fake
=== RUN   TestRouter/Reqs_matched_correctly
=== RUN   TestRouter/Reqs_matched_correctly/POST_/api
=== RUN   TestRouter/Reqs_matched_correctly/POST_/api/
=== RUN   TestRouter/Reqs_matched_correctly/DELETE_/api
=== RUN   TestRouter/Reqs_matched_correctly/GET_/api/fake-id
=== RUN   TestRouter/Reqs_matched_correctly/GET_/api/fake-id/bla
=== RUN   TestRouter/Reqs_matched_correctly/GET_/api/fake-id/stuff/faked-fake
=== RUN   TestRouter/Reqs_execute_correctly
=== RUN   TestRouter/Reqs_execute_correctly/POST_/api_without_auth
=== RUN   TestRouter/Reqs_execute_correctly/POST_/api_with_auth
=== RUN   TestRouter/Reqs_execute_correctly/GET_/api
=== RUN   TestRouter/Overlapping_routes
--- PASS: TestRouter (0.00s)
    --- PASS: TestRouter/Routes_created_correctly (0.00s)
        --- PASS: TestRouter/Routes_created_correctly// (0.00s)
        --- PASS: TestRouter/Routes_created_correctly//:id (0.00s)
        --- PASS: TestRouter/Routes_created_correctly//:id/stuff/:fake (0.00s)
    --- PASS: TestRouter/Reqs_matched_correctly (0.00s)
        --- PASS: TestRouter/Reqs_matched_correctly/POST_/api (0.00s)
        --- PASS: TestRouter/Reqs_matched_correctly/POST_/api/ (0.00s)
        --- PASS: TestRouter/Reqs_matched_correctly/DELETE_/api (0.00s)
        --- PASS: TestRouter/Reqs_matched_correctly/GET_/api/fake-id (0.00s)
        --- PASS: TestRouter/Reqs_matched_correctly/GET_/api/fake-id/bla (0.00s)
        --- PASS: TestRouter/Reqs_matched_correctly/GET_/api/fake-id/stuff/faked-fake (0.00s)
    --- PASS: TestRouter/Reqs_execute_correctly (0.00s)
        --- PASS: TestRouter/Reqs_execute_correctly/POST_/api_without_auth (0.00s)
        --- PASS: TestRouter/Reqs_execute_correctly/POST_/api_with_auth (0.00s)
        --- PASS: TestRouter/Reqs_execute_correctly/GET_/api (0.00s)
    --- PASS: TestRouter/Overlapping_routes (0.00s)
PASS
ok  	github.com/seantcanavan/lambda_jwt_router/lambda_router	0.004s
?   	github.com/seantcanavan/lambda_jwt_router/lambda_util	[no test files]
```
