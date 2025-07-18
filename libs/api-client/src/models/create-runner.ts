/* tslint:disable */

/**
 * Daytona
 * Daytona AI platform API Docs
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@daytona.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 *
 * @export
 * @interface CreateRunner
 */
export interface CreateRunner {
  /**
   *
   * @type {string}
   * @memberof CreateRunner
   */
  domain: string
  /**
   *
   * @type {string}
   * @memberof CreateRunner
   */
  apiUrl: string
  /**
   *
   * @type {string}
   * @memberof CreateRunner
   */
  apiKey: string
  /**
   *
   * @type {number}
   * @memberof CreateRunner
   */
  cpu: number
  /**
   *
   * @type {number}
   * @memberof CreateRunner
   */
  memory: number
  /**
   *
   * @type {number}
   * @memberof CreateRunner
   */
  disk: number
  /**
   *
   * @type {number}
   * @memberof CreateRunner
   */
  gpu: number
  /**
   *
   * @type {string}
   * @memberof CreateRunner
   */
  gpuType: string
  /**
   *
   * @type {string}
   * @memberof CreateRunner
   */
  class: CreateRunnerClassEnum
  /**
   *
   * @type {number}
   * @memberof CreateRunner
   */
  capacity: number
  /**
   *
   * @type {string}
   * @memberof CreateRunner
   */
  region: CreateRunnerRegionEnum
}

export const CreateRunnerClassEnum = {
  SMALL: 'small',
  MEDIUM: 'medium',
  LARGE: 'large',
} as const

export type CreateRunnerClassEnum = (typeof CreateRunnerClassEnum)[keyof typeof CreateRunnerClassEnum]
export const CreateRunnerRegionEnum = {
  EU: 'eu',
  US: 'us',
  ASIA: 'asia',
} as const

export type CreateRunnerRegionEnum = (typeof CreateRunnerRegionEnum)[keyof typeof CreateRunnerRegionEnum]
