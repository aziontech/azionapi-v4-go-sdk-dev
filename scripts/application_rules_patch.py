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
    
    # Define all the behavior request types we need to handle
    behavior_request_types = [
        'ApplicationRequestPhaseBehaviorAddHeaderRequest',
        'ApplicationRequestPhaseBehaviorAddRequestCookieRequest',
        'ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest',
        'ApplicationRequestPhaseBehaviorFilterHeaderRequest',
        'ApplicationRequestPhaseBehaviorFilterRequestCookieRequest',
        'ApplicationRequestPhaseBehaviorNoArgsRequest',
        'ApplicationRequestPhaseBehaviorRewriteRequestRequest',
        'ApplicationRequestPhaseBehaviorRunFunctionRequest',
        'ApplicationRequestPhaseBehaviorSetCachePolicyRequest',
        'ApplicationRequestPhaseBehaviorSetConnectorRequest',
        'ApplicationRequestPhaseBehaviorSetOriginRequest',
        'ApplicationRequestPhaseBehaviorStringRequest'
    ]
    
    # Check if the ApplicationRuleEngineRequestPhaseBehaviorsRequest schema exists
    if 'ApplicationRuleEngineRequestPhaseBehaviorsRequest' not in data['components']['schemas']:
        print("⚠️ Schema ApplicationRuleEngineRequestPhaseBehaviorsRequest not found in YAML, skipping request phase behaviors modification")
        return
    
    # Check if the behavior request types exist
    missing_schemas = []
    for schema in behavior_request_types:
        if schema not in data['components']['schemas']:
            missing_schemas.append(schema)
    
    if missing_schemas:
        print(f"⚠️ The following schemas are missing: {', '.join(missing_schemas)}")
        print("Continuing with available schemas...")
    
    # Create properties for the ApplicationRuleEngineRequestPhaseBehaviorsRequest schema
    properties = {}
    for behavior_type in behavior_request_types:
        if behavior_type in data['components']['schemas']:
            properties[behavior_type] = {
                "$ref": f"#/components/schemas/{behavior_type}"
            }
    
    # Update the ApplicationRuleEngineRequestPhaseBehaviorsRequest schema
    data['components']['schemas']['ApplicationRuleEngineRequestPhaseBehaviorsRequest'] = {
        "type": "object",
        "properties": properties
    }
    
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
    
    # Define all the behavior request types we need to handle
    behavior_request_types = [
        'ApplicationResponsePhaseBehaviorAddHeaderRequest',
        'ApplicationResponsePhaseBehaviorAddResponseCookieRequest',
        'ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest',
        'ApplicationResponsePhaseBehaviorFilterHeaderRequest',
        'ApplicationResponsePhaseBehaviorFilterResponseCookieRequest',
        'ApplicationResponsePhaseBehaviorNoArgsRequest',
        'ApplicationResponsePhaseBehaviorRedirectRequest',
        'ApplicationResponsePhaseBehaviorRewriteResponseRequest',
        'ApplicationResponsePhaseBehaviorRunFunctionRequest',
        'ApplicationResponsePhaseBehaviorSetCachePolicyRequest',
        'ApplicationResponsePhaseBehaviorStringRequest'
    ]
    
    # Check if the ApplicationRuleEngineResponsePhaseBehaviorsRequest schema exists
    if 'ApplicationRuleEngineResponsePhaseBehaviorsRequest' not in data['components']['schemas']:
        print("⚠️ Schema ApplicationRuleEngineResponsePhaseBehaviorsRequest not found in YAML, skipping response phase behaviors modification")
        return
    
    # Check if the behavior request types exist
    missing_schemas = []
    for schema in behavior_request_types:
        if schema not in data['components']['schemas']:
            missing_schemas.append(schema)
    
    if missing_schemas:
        print(f"⚠️ The following schemas are missing: {', '.join(missing_schemas)}")
        print("Continuing with available schemas...")
    
    # Create properties for the ApplicationRuleEngineResponsePhaseBehaviorsRequest schema
    properties = {}
    for behavior_type in behavior_request_types:
        if behavior_type in data['components']['schemas']:
            properties[behavior_type] = {
                "$ref": f"#/components/schemas/{behavior_type}"
            }
    
    # Update the ApplicationRuleEngineResponsePhaseBehaviorsRequest schema
    data['components']['schemas']['ApplicationRuleEngineResponsePhaseBehaviorsRequest'] = {
        "type": "object",
        "properties": properties
    }

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
