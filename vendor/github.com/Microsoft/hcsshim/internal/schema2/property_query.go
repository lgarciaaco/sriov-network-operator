/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

//   By default the basic properties will be returned. This query provides a way to  request specific properties.
type PropertyQuery struct {
	PropertyTypes []PropertyType `json:"PropertyTypes,omitempty"`
}