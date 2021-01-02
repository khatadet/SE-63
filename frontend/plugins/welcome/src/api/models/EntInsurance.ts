/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntInsuranceEdges,
    EntInsuranceEdgesFromJSON,
    EntInsuranceEdgesFromJSONTyped,
    EntInsuranceEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntInsurance
 */
export interface EntInsurance {
    /**
     * Insurancecompany holds the value of the "Insurancecompany" field.
     * @type {string}
     * @memberof EntInsurance
     */
    insurancecompany?: string;
    /**
     * 
     * @type {EntInsuranceEdges}
     * @memberof EntInsurance
     */
    edges?: EntInsuranceEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntInsurance
     */
    id?: number;
}

export function EntInsuranceFromJSON(json: any): EntInsurance {
    return EntInsuranceFromJSONTyped(json, false);
}

export function EntInsuranceFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInsurance {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'insurancecompany': !exists(json, 'Insurancecompany') ? undefined : json['Insurancecompany'],
        'edges': !exists(json, 'edges') ? undefined : EntInsuranceEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntInsuranceToJSON(value?: EntInsurance | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Insurancecompany': value.insurancecompany,
        'edges': EntInsuranceEdgesToJSON(value.edges),
        'id': value.id,
    };
}

