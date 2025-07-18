/* tslint:disable */

/**
 * Daytona Runner API
 * Daytona Runner API
 *
 * The version of the OpenAPI document: v0.0.0-dev
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 *
 * @export
 * @enum {string}
 */

export const EnumsBackupState = {
  BackupStateNone: 'NONE',
  BackupStatePending: 'PENDING',
  BackupStateInProgress: 'IN_PROGRESS',
  BackupStateCompleted: 'COMPLETED',
  BackupStateFailed: 'FAILED',
} as const

export type EnumsBackupState = (typeof EnumsBackupState)[keyof typeof EnumsBackupState]
