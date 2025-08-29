#!/usr/bin/env python3
import sys
import yaml
import re

def replace_strings_in_yaml(content):
    """
    Replace specific strings in the YAML content.
    """
    replacements = {
        # Old replacements
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineNoArgsRequest": "ApplicationRequestPhaseBehaviorWithoutArgsRequest",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineStringRequest": "ApplicationRequestPhaseBehaviorWithArgsRequest",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest": "ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineNoArgs": "ApplicationRequestPhaseBehaviorWithoutArgs",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineString": "ApplicationRequestPhaseBehaviorWithArgs",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroups": "ApplicationRequestPhaseBehaviorCaptureMatchGroups",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseNoArgsRequest": "ApplicationResponsePhaseBehaviorWithoutArgsRequest",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseStringRequest": "ApplicationResponsePhaseBehaviorWithArgsRequest",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest":  "ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseNoArgs": "ApplicationResponsePhaseBehaviorWithoutArgs",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseString": "ApplicationResponsePhaseBehaviorWithArgs",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroups":  "ApplicationResponsePhaseBehaviorCaptureMatchGroups",
        
        # New replacements with updated naming convention (removed Edge prefix)
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest": "ApplicationRequestPhaseBehaviorAddHeaderRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddRequestCookieRequest": "ApplicationRequestPhaseBehaviorAddRequestCookieRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineCaptureMatchGroupsRequest": "ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterHeaderRequest": "ApplicationRequestPhaseBehaviorFilterHeaderRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterRequestCookieRequest": "ApplicationRequestPhaseBehaviorFilterRequestCookieRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineNoArgsRequest": "ApplicationRequestPhaseBehaviorNoArgsRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRewriteRequestRequest": "ApplicationRequestPhaseBehaviorRewriteRequestRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRunFunctionRequest": "ApplicationRequestPhaseBehaviorRunFunctionRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetCachePolicyRequest": "ApplicationRequestPhaseBehaviorSetCachePolicyRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetConnectorRequest": "ApplicationRequestPhaseBehaviorSetConnectorRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetOriginRequest": "ApplicationRequestPhaseBehaviorSetOriginRequest",
        "ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineStringRequest": "ApplicationRequestPhaseBehaviorStringRequest"
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
        'ApplicationRuleEngineStringAttributes'
    ]
    
    for schema in required_schemas:
        if schema not in data['components']['schemas']:
            print(f"⚠️ Schema {schema} not found in YAML, skipping request phase behaviors modification")
            return
    
    allowed_refs = [
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithoutArgs'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithArgs'},
        # {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorCaptureMatchGroups'}
    ]
    
    allowed_refs_request = [
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithoutArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorWithArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest'}
    ]

    # Update oneOfs
    data['components']['schemas']['ApplicationRuleEngineRequestPhaseBehaviors']['oneOf'] = allowed_refs
    data['components']['schemas']['ApplicationRuleEngineRequestPhaseBehaviorsRequest']['oneOf'] = allowed_refs_request
    
    # Remove additionalProperties from withoutArgs struct to prevent problems
    if data['components']['schemas']['ApplicationRequestPhaseBehaviorWithoutArgs'].get('additionalProperties', False):
        data['components']['schemas']['ApplicationRequestPhaseBehaviorWithoutArgs']['additionalProperties'] = False
    
    # Adapt allOf behavior schema
    data['components']['schemas']['ApplicationRequestPhaseBehaviorWithoutArgs']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['ApplicationRequestPhaseBehaviorWithArgs']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineString'}]
    data['components']['schemas']['ApplicationRequestPhaseBehaviorCaptureMatchGroups']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineCaptureMatchGroupsRequest'}]
    data['components']['schemas']['ApplicationRequestPhaseBehaviorWithoutArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['ApplicationRequestPhaseBehaviorWithArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineString'}]
    data['components']['schemas']['ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineCaptureMatchGroupsRequest'}]
    
    # Adapt value field validation
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
    
    for schema in required_schemas:
        if schema not in data['components']['schemas']:
            print(f"⚠️ Schema {schema} not found in YAML, skipping response phase behaviors modification")
            return
    
    allowed_refs = [
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithoutArgs'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithArgs'},
        # {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorCaptureMatchGroups'}
    ]
    
    allowed_refs_request = [
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithoutArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorWithArgsRequest'},
        {'$ref': '#/components/schemas/ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest'}
    ]

    # Update oneOfs
    data['components']['schemas']['ApplicationRuleEngineResponsePhaseBehaviors']['oneOf'] = allowed_refs
    data['components']['schemas']['ApplicationRuleEngineResponsePhaseBehaviorsRequest']['oneOf'] = allowed_refs_request
    
    # Remove additionalProperties from withoutArgs struct to prevent problems
    if data['components']['schemas']['ApplicationResponsePhaseBehaviorWithoutArgs'].get('additionalProperties', False):
        data['components']['schemas']['ApplicationResponsePhaseBehaviorWithoutArgs']['additionalProperties'] = False
    
    # Adapt allOf behavior schema
    data['components']['schemas']['ApplicationResponsePhaseBehaviorWithoutArgs']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['ApplicationResponsePhaseBehaviorWithArgs']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineString'}]
    data['components']['schemas']['ApplicationResponsePhaseBehaviorCaptureMatchGroups']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineCaptureMatchGroups'}]
    data['components']['schemas']['ApplicationResponsePhaseBehaviorWithoutArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['ApplicationResponsePhaseBehaviorWithArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineString'}]
    data['components']['schemas']['ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest']['allOf'] = [{'$ref': '#/components/schemas/ApplicationRuleEngineCaptureMatchGroups'}]


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

