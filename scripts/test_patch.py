#!/usr/bin/env python3
import yaml
import sys
from application_rules_patch import replace_strings_in_yaml

# Test YAML content with ApplicationRuleEngine structures
test_yaml = """
components:
  schemas:
    ApplicationRuleEngineRequestPhaseBehaviorsRequest:
      type: object
      properties:
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddRequestCookieRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddRequestCookieRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineCaptureMatchGroupsRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineCaptureMatchGroupsRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterHeaderRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterHeaderRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterRequestCookieRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineFilterRequestCookieRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineNoArgsRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineNoArgsRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRewriteRequestRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRewriteRequestRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRunFunctionRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineRunFunctionRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetCachePolicyRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetCachePolicyRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetConnectorRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetConnectorRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetOriginRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineSetOriginRequest'
        ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineStringRequest:
          $ref: '#/components/schemas/ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineStringRequest'
"""

def test_replacements():
    """Test the string replacements in the YAML content."""
    print("Testing string replacements...")
    
    # Apply the replacements
    modified_content = replace_strings_in_yaml(test_yaml)
    
    # Check if the replacements were applied correctly
    expected_replacements = [
        "ApplicationRequestPhaseBehaviorWithoutArgsRequest",
        "ApplicationRequestPhaseBehaviorWithArgsRequest",
        "ApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest",
        "ApplicationRequestPhaseBehaviorAddHeaderRequest",
        "ApplicationRequestPhaseBehaviorAddRequestCookieRequest",
        "ApplicationRequestPhaseBehaviorFilterHeaderRequest",
        "ApplicationRequestPhaseBehaviorFilterRequestCookieRequest",
        "ApplicationRequestPhaseBehaviorRewriteRequestRequest",
        "ApplicationRequestPhaseBehaviorRunFunctionRequest",
        "ApplicationRequestPhaseBehaviorSetCachePolicyRequest",
        "ApplicationRequestPhaseBehaviorSetConnectorRequest",
        "ApplicationRequestPhaseBehaviorSetOriginRequest"
    ]
    
    success = True
    for replacement in expected_replacements:
        if replacement not in modified_content:
            print(f"❌ Failed: {replacement} not found in modified content")
            success = False
    
    if success:
        print("✅ All replacements were applied correctly")
    
    # Print the modified content for inspection
    print("\nModified YAML content:")
    print(modified_content)

if __name__ == "__main__":
    test_replacements()
