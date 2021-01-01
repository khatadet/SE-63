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
/**
 * 
 * @export
 * @interface ControllersPatientrights
 */
export interface ControllersPatientrights {
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    abilitypatientrights?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    insurance?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    medicalrecordstaff?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    patientrecord?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    patientrightstype?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrights
     */
    permissionDate?: string;
}

export function ControllersPatientrightsFromJSON(json: any): ControllersPatientrights {
    return ControllersPatientrightsFromJSONTyped(json, false);
}

export function ControllersPatientrightsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPatientrights {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'abilitypatientrights': !exists(json, 'abilitypatientrights') ? undefined : json['abilitypatientrights'],
        'insurance': !exists(json, 'insurance') ? undefined : json['insurance'],
        'medicalrecordstaff': !exists(json, 'medicalrecordstaff') ? undefined : json['medicalrecordstaff'],
        'patientrecord': !exists(json, 'patientrecord') ? undefined : json['patientrecord'],
        'patientrightstype': !exists(json, 'patientrightstype') ? undefined : json['patientrightstype'],
        'permissionDate': !exists(json, 'permissionDate') ? undefined : json['permissionDate'],
    };
}

export function ControllersPatientrightsToJSON(value?: ControllersPatientrights | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'abilitypatientrights': value.abilitypatientrights,
        'insurance': value.insurance,
        'medicalrecordstaff': value.medicalrecordstaff,
        'patientrecord': value.patientrecord,
        'patientrightstype': value.patientrightstype,
        'permissionDate': value.permissionDate,
    };
}


