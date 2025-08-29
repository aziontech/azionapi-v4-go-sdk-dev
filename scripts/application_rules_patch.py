#!/usr/bin/env python3
import sys
import yaml
import re

def replace_strings_in_yaml(content):
    """
    Replace specific strings in the YAML content.
    """
    replacements = {
        # Request Phase Behaviors
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineNoArgsRequest": "ApplicationRequestPhaseBehaviorWithoutArgsRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineStringRequest": "ApplicationRequestPhaseBehaviorWithArgsRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineCaptureMatchGroupsRequest": "ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest": "ApplicationRequestPhaseBehaviorAddHeaderRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddRequestCookieRequest": "ApplicationRequestPhaseBehaviorAddRequestCookieRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterHeaderRequest": "ApplicationRequestPhaseBehaviorFilterHeaderRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterRequestCookieRequest": "ApplicationRequestPhaseBehaviorFilterRequestCookieRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRewriteRequestRequest": "ApplicationRequestPhaseBehaviorRewriteRequestRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRunFunctionRequest": "ApplicationRequestPhaseBehaviorRunFunctionRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetCachePolicyRequest": "ApplicationRequestPhaseBehaviorSetCachePolicyRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetConnectorRequest": "ApplicationRequestPhaseBehaviorSetConnectorRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetOriginRequest": "ApplicationRequestPhaseBehaviorSetOriginRequest",
        
        # Request Phase Behaviors (non-request versions)
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineNoArgs": "ApplicationRequestPhaseBehaviorWithoutArgs",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineString": "ApplicationRequestPhaseBehaviorWithArgs",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineCaptureMatchGroups": "ApplicationRequestPhaseBehaviorCaptureMatchGroups",
        
        # Response Phase Behaviors
        "ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineResponseNoArgsRequest": "ApplicationResponsePhaseBehaviorWithoutArgsRequest",
        "ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineResponseStringRequest": "ApplicationResponsePhaseBehaviorWithArgsRequest",
        "ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineCaptureMatchGroupsRequest":  "ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest",
        
        # Response Phase Behaviors (non-request versions)
        "ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineResponseNoArgs": "ApplicationResponsePhaseBehaviorWithoutArgs",
        "ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineResponseString": "ApplicationResponsePhaseBehaviorWithArgs",
        "ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineCaptureMatchGroups":  "ApplicationResponsePhaseBehaviorCaptureMatchGroups"
        
    }
    
    # Perform string replacements
    for old_str, new_str in replacements.items():
        content = content.replace(old_str, new_str)
    
    print("✅ Successfully replaced strings in YAML")
    return content


def modify_request_phase_behaviors(data):
    """
    Modify request phase behaviors
    """
    
    # Check if the required schemas exist before proceeding
    if 'components' not in data or 'schemas' not in data['components']:
        print("⚠️ No components.schemas found in YAML, skipping request phase behaviors modification")
        return
        
    # Check if the specific schemas we need to modify exist
    required_schemas = [
        'ApplicationRuleEngineRequestPhaseBehaviors',
        'ApplicationRuleEngineRequestPhaseBehaviorsRequest',
        'ApplicationRequestPhaseBehaviorWithoutArgs',
        'ApplicationRequestPhaseBehaviorWithArgs',
        'ApplicationRequestPhaseBehaviorCaptureMatchGroups',
        'ApplicationRequestPhaseBehaviorWithoutArgsRequest',
        'ApplicationRequestPhaseBehaviorWithArgsRequest',
        'ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest',
        'ApplicationRequestPhaseBehaviorAddHeaderRequest',
        'ApplicationRequestPhaseBehaviorAddRequestCookieRequest',
        'ApplicationRequestPhaseBehaviorFilterHeaderRequest',
        'ApplicationRequestPhaseBehaviorFilterRequestCookieRequest',
        'ApplicationRequestPhaseBehaviorRewriteRequestRequest',
        'ApplicationRequestPhaseBehaviorRunFunctionRequest',
        'ApplicationRequestPhaseBehaviorSetCachePolicyRequest',
        'ApplicationRequestPhaseBehaviorSetConnectorRequest',
        'ApplicationRequestPhaseBehaviorSetOriginRequest',
        'ApplicationRuleEngineStringAttributes'
    ]
    
    # Check for required schemas, but don't fail if some are missing
    missing_schemas = []
    for schema in required_schemas:
        if schema not in data['components']['schemas']:
            missing_schemas.append(schema)
    
    if missing_schemas:
        print(f"⚠️ Some schemas are missing: {', '.join(missing_schemas)}")
        print("⚠️ Will attempt to continue with available schemas")
    
    # Define all possible behavior references
    behavior_refs = [
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithoutArgs'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithArgs'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorCaptureMatchGroups'}
    ]
    
    behavior_request_refs = [
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithoutArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorAddHeaderRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorAddRequestCookieRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorFilterHeaderRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorFilterRequestCookieRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorRewriteRequestRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorRunFunctionRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorSetCachePolicyRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorSetConnectorRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorSetOriginRequest'}
    ]

    # Filter out references to missing schemas
    allowed_refs = [ref for ref in behavior_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]
    allowed_refs_request = [ref for ref in behavior_request_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]

    # Update oneOfs if the schemas exist
    if 'ApplicationRuleEngineRequestPhaseBehaviors' in data['components']['schemas']:
        data['components']['schemas']['ApplicationRuleEngineRequestPhaseBehaviors']['oneOf'] = allowed_refs
    
    if 'ApplicationRuleEngineRequestPhaseBehaviorsRequest' in data['components']['schemas']:
        data['components']['schemas']['ApplicationRuleEngineRequestPhaseBehaviorsRequest']['oneOf'] = allowed_refs_request
    
    # Remove additionalProperties from withoutArgs struct to prevent problems
    if 'ApplicationRequestPhaseBehaviorWithoutArgs' in data['components']['schemas'] and \
       data['components']['schemas']['ApplicationRequestPhaseBehaviorWithoutArgs'].get('additionalProperties', False):
        data['components']['schemas']['ApplicationRequestPhaseBehaviorWithoutArgs']['additionalProperties'] = False
    
    # Adapt allOf behavior schema - only if the schemas exist
    schema_mappings = {
        'ApplicationRequestPhaseBehaviorWithoutArgs': 'ApplicationRuleEngineNoArgs',
        'ApplicationRequestPhaseBehaviorWithArgs': 'ApplicationRuleEngineString',
        'ApplicationRequestPhaseBehaviorCaptureMatchGroups': 'ApplicationRuleEngineCaptureMatchGroupsRequest',
        'ApplicationRequestPhaseBehaviorWithoutArgsRequest': 'ApplicationRuleEngineNoArgs',
        'ApplicationRequestPhaseBehaviorWithArgsRequest': 'ApplicationRuleEngineString',
        'ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest': 'ApplicationRuleEngineCaptureMatchGroupsRequest'
    }
    
    for target_schema, ref_schema in schema_mappings.items():
        if target_schema in data['components']['schemas'] and ref_schema in data['components']['schemas']:
            data['components']['schemas'][target_schema]['allOf'] = [{'$ref': f'#/components/schemas/{ref_schema}'}]
    
    # Adapt value field validation if the schema exists
    if 'ApplicationRuleEngineStringAttributes' in data['components']['schemas']:
        data['components']['schemas']['ApplicationRuleEngineStringAttributes']['properties']['value'] = {
            "oneOf": [
                { "type": "string" },
                { "type": "integer", "format": "int64" }
            ]
        }
        
        # Remove Required Fields
        if 'required' in data['components']['schemas']['ApplicationRuleEngineStringAttributes']:
            del data['components']['schemas']['ApplicationRuleEngineStringAttributes']['required']

    print("✅ Successfully modified Request Phase Behaviors")


def modify_response_phase_behaviors(data):
    """
    Modify response phase behaviors
    """
    
    # Check if the required schemas exist before proceeding
    if 'components' not in data or 'schemas' not in data['components']:
        print("⚠️ No components.schemas found in YAML, skipping response phase behaviors modification")
        return
        
    # Check if the specific schemas we need to modify exist
    required_schemas = [
        'ApplicationRuleEngineResponsePhaseBehaviors',
        'ApplicationRuleEngineResponsePhaseBehaviorsRequest',
        'ApplicationResponsePhaseBehaviorWithoutArgs',
        'ApplicationResponsePhaseBehaviorWithArgs',
        'ApplicationResponsePhaseBehaviorCaptureMatchGroups',
        'ApplicationResponsePhaseBehaviorWithoutArgsRequest',
        'ApplicationResponsePhaseBehaviorWithArgsRequest',
        'ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest'
    ]
    
    # Check for required schemas, but don't fail if some are missing
    missing_schemas = []
    for schema in required_schemas:
        if schema not in data['components']['schemas']:
            missing_schemas.append(schema)
    
    if missing_schemas:
        print(f"⚠️ Some response schemas are missing: {', '.join(missing_schemas)}")
        print("⚠️ Will attempt to continue with available schemas")
    
    # Define all possible behavior references
    behavior_refs = [
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithoutArgs'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithArgs'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorCaptureMatchGroups'}
    ]
    
    behavior_request_refs = [
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithoutArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest'}
    ]

    # Filter out references to missing schemas
    allowed_refs = [ref for ref in behavior_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]
    allowed_refs_request = [ref for ref in behavior_request_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]

    # Update oneOfs if the schemas exist
    if 'ApplicationRuleEngineResponsePhaseBehaviors' in data['components']['schemas']:
        data['components']['schemas']['ApplicationRuleEngineResponsePhaseBehaviors']['oneOf'] = allowed_refs
    
    if 'ApplicationRuleEngineResponsePhaseBehaviorsRequest' in data['components']['schemas']:
        data['components']['schemas']['ApplicationRuleEngineResponsePhaseBehaviorsRequest']['oneOf'] = allowed_refs_request
    
    # Remove additionalProperties from withoutArgs struct to prevent problems
    if 'ApplicationResponsePhaseBehaviorWithoutArgs' in data['components']['schemas'] and \
       data['components']['schemas']['ApplicationResponsePhaseBehaviorWithoutArgs'].get('additionalProperties', False):
        data['components']['schemas']['ApplicationResponsePhaseBehaviorWithoutArgs']['additionalProperties'] = False
    
    # Adapt allOf behavior schema - only if the schemas exist
    schema_mappings = {
        'ApplicationResponsePhaseBehaviorWithoutArgs': 'ApplicationRuleEngineNoArgs',
        'ApplicationResponsePhaseBehaviorWithArgs': 'ApplicationRuleEngineString',
        'ApplicationResponsePhaseBehaviorCaptureMatchGroups': 'ApplicationRuleEngineCaptureMatchGroups',
        'ApplicationResponsePhaseBehaviorWithoutArgsRequest': 'ApplicationRuleEngineNoArgs',
        'ApplicationResponsePhaseBehaviorWithArgsRequest': 'ApplicationRuleEngineString',
        'ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest': 'ApplicationRuleEngineCaptureMatchGroups'
    }
    
    for target_schema, ref_schema in schema_mappings.items():
        if target_schema in data['components']['schemas'] and ref_schema in data['components']['schemas']:
            data['components']['schemas'][target_schema]['allOf'] = [{'$ref': f'#/components/schemas/{ref_schema}'}]

    print("✅ Successfully modified Response Phase Behaviors")

    
def process_yaml_file(file_path):
    """
    Process a YAML file by loading it, replacing strings, and modifying component schema.
    """
    try:
        # Read the file
        with open(file_path, 'r') as file:
            yaml_content = file.read()
        
        # Replace strings
        modified_content = replace_strings_in_yaml(yaml_content)
        
        # Load YAML to modify component schema
        data = yaml.safe_load(modified_content)
        
        # Modify behaviors
        modify_request_phase_behaviors(data)
        modify_response_phase_behaviors(data)
        
        # Write back to file
        with open(file_path, 'w') as file:
            yaml.dump(data, file, default_flow_style=False, sort_keys=False)
        
        print(f"✅ Successfully processed {file_path}")
        
    except Exception as e:
        print(f"❌ Error processing {file_path}: {str(e)}")
        sys.exit(1)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python replace_strings.py <path-to-yaml-file>")
        sys.exit(1)
    
    yaml_file_path = sys.argv[1]
    process_yaml_file(yaml_file_path)
