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
        'RequestPhaseBehavior',
        'RequestPhaseBehaviorRequest',
        'BehaviorNoArgs',
        'BehaviorWithArgs',
        'BehaviorCapture',
        'BehaviorAttributes'
    ]
    
    # Check for required schemas, but don't fail if some are missing
    missing_schemas = []
    for schema in required_schemas:
        if schema not in data['components']['schemas']:
            missing_schemas.append(schema)
    
    if missing_schemas:
        print(f"⚠️ Some schemas are missing: {', '.join(missing_schemas)}")
        print("⚠️ Will attempt to continue with available schemas")

    # Create or update BehaviorAttributes schema FIRST to handle both string and integer values
    if 'BehaviorAttributes' not in data['components']['schemas']:
        data['components']['schemas']['BehaviorAttributes'] = {}
    
    data['components']['schemas']['BehaviorAttributes']['type'] = 'object'
    data['components']['schemas']['BehaviorAttributes']['properties'] = {
        'value': {
            'oneOf': [
                {'type': 'string'},
                {'type': 'integer', 'format': 'int64'}
            ]
        }
    }
    data['components']['schemas']['BehaviorAttributes']['additionalProperties'] = False
    
    if 'BehaviorWithArgs' not in data['components']['schemas']:
        data['components']['schemas']['BehaviorWithArgs'] = {}
    
    data['components']['schemas']['BehaviorWithArgs']['type'] = 'object'
    data['components']['schemas']['BehaviorWithArgs']['properties'] = {
        'type': {'type': 'string'},
        'attributes': {'$ref': '#/components/schemas/BehaviorAttributes'}
    }
    data['components']['schemas']['BehaviorWithArgs']['required'] = ['type', 'attributes']
    data['components']['schemas']['BehaviorWithArgs']['additionalProperties'] = False

    # Define all possible behavior references (matching new oneOf structure)
    behavior_refs = [
        {'$ref': '#/components/schemas/BehaviorNoArgs'},
        {'$ref': '#/components/schemas/BehaviorWithArgs'},
        {'$ref': '#/components/schemas/BehaviorCapture'}
    ]
    
    # For Request behaviors, we consolidate to the 3 main types
    behavior_request_refs = [
        {'$ref': '#/components/schemas/BehaviorNoArgs'},
        {'$ref': '#/components/schemas/BehaviorWithArgs'},
        {'$ref': '#/components/schemas/BehaviorCapture'}
    ]

    # Filter out references to missing schemas (now BehaviorWithArgs exists!)
    allowed_refs = [ref for ref in behavior_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]
    allowed_refs_request = [ref for ref in behavior_request_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]

    # Update oneOfs if the schemas exist
    if 'RequestPhaseBehavior' in data['components']['schemas']:
        data['components']['schemas']['RequestPhaseBehavior']['oneOf'] = allowed_refs
    
    if 'RequestPhaseBehaviorRequest' in data['components']['schemas']:
        # Merge BehaviorString and BehaviorInteger into BehaviorWithArgs
        # This creates a unified structure that handles both string and integer values
        desired_request_oneof = [
            {'$ref': '#/components/schemas/BehaviorNoArgs'},
            {'$ref': '#/components/schemas/BehaviorWithArgs'},
            {'$ref': '#/components/schemas/BehaviorCapture'},
        ]
        data['components']['schemas']['RequestPhaseBehaviorRequest']['oneOf'] = [
            ref for ref in desired_request_oneof
            if ref['$ref'].split('/')[-1] in data['components']['schemas']
        ]

        disc = data['components']['schemas']['RequestPhaseBehaviorRequest'].get('discriminator')
        if isinstance(disc, dict) and isinstance(disc.get('mapping'), dict):
            without_args_types = {
                'bypass_cache',
                'deliver',
                'deny',
                'enable_gzip',
                'forward_cookies',
                'no_content',
                'redirect_http_to_https',
            }
            new_mapping = {}
            for k, v in disc['mapping'].items():
                if not isinstance(v, str):
                    continue
                if k == 'capture_match_groups':
                    new_mapping[k] = '#/components/schemas/BehaviorCapture'
                elif k in without_args_types:
                    new_mapping[k] = '#/components/schemas/BehaviorNoArgs'
                else:
                    # Both string and integer behaviors map to BehaviorWithArgs
                    new_mapping[k] = '#/components/schemas/BehaviorWithArgs'
            disc['mapping'] = new_mapping
    
    # Remove additionalProperties from BehaviorNoArgs struct to prevent problems
    if 'BehaviorNoArgs' in data['components']['schemas'] and \
       data['components']['schemas']['BehaviorNoArgs'].get('additionalProperties', False):
        data['components']['schemas']['BehaviorNoArgs']['additionalProperties'] = False
    
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
        'ResponsePhaseBehavior',
        'ResponsePhaseBehaviorRequest',
        'BehaviorNoArgs',
        'BehaviorWithArgs',
        'BehaviorCapture',
        'BehaviorAttributes'
    ]
    
    # Check for required schemas, but don't fail if some are missing
    missing_schemas = []
    for schema in required_schemas:
        if schema not in data['components']['schemas']:
            missing_schemas.append(schema)
    
    if missing_schemas:
        print(f"⚠️ Some response schemas are missing: {', '.join(missing_schemas)}")
        print("⚠️ Will attempt to continue with available schemas")
    
    # Define all possible behavior references (matching new oneOf structure)
    behavior_refs = [
        {'$ref': '#/components/schemas/BehaviorNoArgs'},
        {'$ref': '#/components/schemas/BehaviorWithArgs'},
        {'$ref': '#/components/schemas/BehaviorCapture'}
    ]
    
    # For Response behaviors, we consolidate to the 3 main types
    behavior_request_refs = [
        {'$ref': '#/components/schemas/BehaviorNoArgs'},
        {'$ref': '#/components/schemas/BehaviorWithArgs'},
        {'$ref': '#/components/schemas/BehaviorCapture'}
    ]

    # Filter out references to missing schemas
    allowed_refs = [ref for ref in behavior_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]
    allowed_refs_request = [ref for ref in behavior_request_refs if ref['$ref'].split('/')[-1] in data['components']['schemas']]

    # Update oneOfs if the schemas exist
    if 'ResponsePhaseBehavior' in data['components']['schemas']:
        data['components']['schemas']['ResponsePhaseBehavior']['oneOf'] = allowed_refs

        disc = data['components']['schemas']['ResponsePhaseBehavior'].get('discriminator')
        if isinstance(disc, dict) and isinstance(disc.get('mapping'), dict):
            without_args_types = {
                'enable_gzip',
                'deliver',
            }

            new_mapping = {}
            for k, v in disc['mapping'].items():
                if not isinstance(v, str):
                    continue
                if k == 'capture_match_groups':
                    new_mapping[k] = '#/components/schemas/BehaviorCapture'
                elif k in without_args_types:
                    new_mapping[k] = '#/components/schemas/BehaviorNoArgs'
                else:
                    new_mapping[k] = '#/components/schemas/BehaviorWithArgs'
            disc['mapping'] = new_mapping

    if 'ResponsePhaseBehaviorRequest' in data['components']['schemas']:
        # First, update the discriminator mapping to use BehaviorWithArgs
        disc = data['components']['schemas']['ResponsePhaseBehaviorRequest'].get('discriminator')
        if isinstance(disc, dict) and isinstance(disc.get('mapping'), dict):
            new_mapping = {}
            for k, v in disc['mapping'].items():
                if not isinstance(v, str):
                    continue
                if k == 'capture_match_groups':
                    new_mapping[k] = '#/components/schemas/BehaviorCapture'
                elif k in {'enable_gzip', 'deliver'}:
                    new_mapping[k] = '#/components/schemas/BehaviorNoArgs'
                else:
                    # Both string and integer behaviors map to BehaviorWithArgs
                    new_mapping[k] = '#/components/schemas/BehaviorWithArgs'
            disc['mapping'] = new_mapping
        
        # Then update the oneOf to only reference the 3 consolidated schemas
        # This replaces BehaviorString and BehaviorInteger with BehaviorWithArgs
        data['components']['schemas']['ResponsePhaseBehaviorRequest']['oneOf'] = [
            {'$ref': '#/components/schemas/BehaviorNoArgs'},
            {'$ref': '#/components/schemas/BehaviorWithArgs'},
            {'$ref': '#/components/schemas/BehaviorCapture'},
        ]
    
    print("✅ Successfully modified Response Phase Behaviors")

    
def cleanup_merged_schemas(data):
    """
    Remove BehaviorString and BehaviorInteger schemas after both phases are processed.
    These have been merged into BehaviorWithArgs.
    """
    if 'components' not in data or 'schemas' not in data['components']:
        return
    
    # Remove BehaviorInteger schema if it exists (merged into BehaviorWithArgs)
    if 'BehaviorInteger' in data['components']['schemas']:
        del data['components']['schemas']['BehaviorInteger']
        print("✅ Removed BehaviorInteger schema (merged into BehaviorWithArgs)")
    
    # Remove BehaviorString schema if it exists (merged into BehaviorWithArgs)
    if 'BehaviorString' in data['components']['schemas']:
        del data['components']['schemas']['BehaviorString']
        print("✅ Removed BehaviorString schema (merged into BehaviorWithArgs)")


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
        
        # Clean up merged schemas after both phases are processed
        cleanup_merged_schemas(data)
        
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
