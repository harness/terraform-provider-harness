# Central Notifications Test Plan

## Overview
This document outlines a comprehensive test plan for the Central Notification System in the Harness Terraform Provider, covering Channels, Rules, and Templates components.

## Components Under Test

### 1. Central Notification Channels
**File Location:** `internal/service/platform/central_notification_channel/`
- Resource: `harness_platform_central_notification_channel`
- Data Source: `harness_central_notification_channel`

### 2. Central Notification Rules  
**File Location:** `internal/service/platform/central_notification_rule/`
- Resource: `harness_platform_central_notification_rule`
- Data Source: `harness_central_notification_rule`

### 3. Default Notification Template Sets
**File Location:** `internal/service/platform/default_notification_template_set/`
- Resource: `harness_platform_default_notification_template_set`
- Data Source: `harness_default_notification_template_set`

## Test Coverage Analysis

### ✅ COVERED - Central Notification Channels

#### Unit Tests Implemented:
1. **Basic Channel Creation/Update/Delete** (`TestAccResourceCentralNotificationChannel_basic`)
   - EMAIL channel type
   - Basic CRUD operations
   - Import functionality

2. **Slack Channel** (`TestAccResourceCentralNotificationChannel_slack`)
   - SLACK channel type with webhook URLs
   - Import functionality

3. **Multi-Scope Testing** (`TestAccResourceCentralNotificationChannel_projectLevel`)
   - Project-level channel creation
   - Organization and project dependencies

4. **Multiple Channel Types** (`TestAccResourceCentralNotificationChannel_multipleChannelTypes`)
   - MS Teams channels
   - PagerDuty channels

#### Supported Channel Types:
- EMAIL ✅
- SLACK ✅  
- PAGERDUTY ✅
- MSTEAMS ✅
- WEBHOOK (partial)
- DATADOG (partial)

### ✅ COVERED - Central Notification Rules

#### Unit Tests Implemented:
1. **Basic Rule Creation** (`TestAccResourceCentralNotificationRule`)
   - Basic CRUD operations
   - Single condition with PIPELINE entity
   - Import functionality

2. **Multiple Conditions** (`TestAccResourceCentralNotificationRule_multipleConditions`)
   - Pipeline and deployment conditions
   - Multiple notification event configs

3. **Multi-Scope Testing**
   - Organization level (`TestAccResourceCentralNotificationRule_orgLevel`)
   - Account level (`TestAccResourceCentralNotificationRule_accountLevel`)

#### Supported Notification Entities:
- PIPELINE ✅
- DEPLOYMENT ✅ 
- DELEGATE (schema only)
- CONNECTOR (schema only)
- CHAOS_EXPERIMENT (schema only)
- SERVICE_LEVEL_OBJECTIVE (schema only)
- STO_EXEMPTION (schema only)

### ✅ COVERED - Default Notification Template Sets

#### Unit Tests Implemented:
1. **Basic Template Set** (`TestAccResourceDefaultNotificationTemplateSet_basic`)
   - EMAIL channel type with PIPELINE entity
   - Multiple notification events
   - Template reference configuration

2. **Multi-Scope Testing**
   - Project level (`TestAccResourceDefaultNotificationTemplateSet_projectLevel`)
   - Organization level (`TestAccResourceDefaultNotificationTemplateSet_orgLevel`)

3. **Multiple Channel Types** (`TestAccResourceDefaultNotificationTemplateSet_multipleChannelTypes`)
   - SLACK templates
   - MS Teams templates

4. **Multiple Event Configurations** (`TestAccResourceDefaultNotificationTemplateSet_multipleEventConfigurations`)
   - DEPLOYMENT entity
   - Multiple event template configurations

## 🔍 GAPS IDENTIFIED - Missing Test Coverage

### Central Notification Channels

#### Missing Channel Types:
1. **WEBHOOK Channels**
   - Custom webhook URLs with headers
   - API key authentication
   - Delegate execution

2. **DATADOG Channels**
   - Datadog webhook URLs
   - API key configuration

#### Missing Features:
1. **User Groups Integration**
   - Testing user group notifications
   - Cross-scope user group references

2. **Delegate Selectors**
   - Testing delegate selector functionality
   - Execute on delegate scenarios

3. **Custom Headers for Webhooks**
   - Header key-value pairs
   - Authentication headers

4. **Edge Cases**
   - Invalid channel configurations
   - Missing required fields
   - Large payload handling

### Central Notification Rules

#### Missing Event Types:
1. **Entity-Specific Events**
   - DELEGATE events (DELEGATE_DOWN, DELEGATE_UP)
   - CONNECTOR events (CONNECTOR_VALIDATION_FAILED)
   - CHAOS_EXPERIMENT events
   - SERVICE_LEVEL_OBJECTIVE events
   - STO_EXEMPTION events

2. **Complex Event Data**
   - Entity identifiers testing
   - Notification event data validation
   - Custom event parameters

#### Missing Features:
1. **Custom Notification Templates**
   - Template reference with variables
   - Version label management
   - Template variable validation

2. **Multiple Channel References**
   - Rules with multiple notification channels
   - Cross-scope channel references

3. **Complex Conditions**
   - Rules with multiple conditions
   - Condition-specific event configurations

### Default Notification Template Sets

#### Missing Entities:
1. **SERVICE Entity Templates**
   - Service-specific notification events
   - Service deployment notifications

2. **CONNECTOR Entity Templates**
   - Connector validation templates
   - Connector health notifications

3. **DELEGATE Entity Templates**
   - Delegate status templates
   - Delegate connectivity notifications

#### Missing Features:
1. **Template Variables**
   - Variable type validation
   - Complex variable configurations
   - Variable interpolation testing

2. **Multi-Event Complex Scenarios**
   - Templates spanning multiple entities
   - Cross-entity notification workflows

## 📋 RECOMMENDED TEST ADDITIONS

### High Priority

#### 1. Webhook Channel Tests
```go
func TestAccResourceCentralNotificationChannel_webhook(t *testing.T)
func TestAccResourceCentralNotificationChannel_webhookWithHeaders(t *testing.T)
func TestAccResourceCentralNotificationChannel_webhookWithDelegate(t *testing.T)
```

#### 2. DataDog Channel Tests
```go
func TestAccResourceCentralNotificationChannel_datadog(t *testing.T)
func TestAccResourceCentralNotificationChannel_datadogWithApiKey(t *testing.T)
```

#### 3. User Groups Integration Tests
```go
func TestAccResourceCentralNotificationChannel_userGroups(t *testing.T)
func TestAccResourceCentralNotificationRule_userGroupsIntegration(t *testing.T)
```

#### 4. Custom Template Tests
```go
func TestAccResourceCentralNotificationRule_customTemplate(t *testing.T)
func TestAccResourceCentralNotificationRule_templateWithVariables(t *testing.T)
```

#### 5. Multi-Entity Template Tests
```go
func TestAccResourceDefaultNotificationTemplateSet_serviceEntity(t *testing.T)
func TestAccResourceDefaultNotificationTemplateSet_delegateEntity(t *testing.T)
func TestAccResourceDefaultNotificationTemplateSet_connectorEntity(t *testing.T)
```

### Medium Priority

#### 6. Error Handling Tests
```go
func TestAccResourceCentralNotificationChannel_invalidConfiguration(t *testing.T)
func TestAccResourceCentralNotificationRule_invalidChannelRef(t *testing.T)
func TestAccResourceDefaultNotificationTemplateSet_invalidTemplate(t *testing.T)
```

#### 7. Cross-Scope Reference Tests
```go
func TestAccResourceCentralNotificationRule_crossScopeChannelRef(t *testing.T)
func TestAccResourceCentralNotificationChannel_crossScopeUserGroups(t *testing.T)
```

#### 8. Data Source Tests
```go
func TestAccDataSourceCentralNotificationChannel(t *testing.T)
func TestAccDataSourceCentralNotificationRule(t *testing.T)
func TestAccDataSourceDefaultNotificationTemplateSet(t *testing.T)
```

### Low Priority

#### 9. Performance Tests
```go
func TestAccResourceCentralNotificationRule_largeConditionSet(t *testing.T)
func TestAccResourceCentralNotificationChannel_bulkOperations(t *testing.T)
```

#### 10. Edge Case Tests
```go
func TestAccResourceCentralNotificationChannel_maxLimits(t *testing.T)
func TestAccResourceCentralNotificationRule_emptyConditions(t *testing.T)
```

## 🧪 INTEGRATION TEST SCENARIOS

### End-to-End Workflows

#### 1. Complete Notification Pipeline
```
Channel Creation → Rule Creation → Template Set → Trigger Event → Verify Delivery
```

#### 2. Multi-Channel Notification
```
Multiple Channels (Email + Slack + Teams) → Single Rule → Verify All Channels
```

#### 3. Hierarchical Scope Testing
```
Account → Org → Project level resources and inheritance
```

## 📊 TEST EXECUTION METRICS

### Current Test Coverage
- **Central Notification Channels**: ~70% (4/6 channel types)
- **Central Notification Rules**: ~60% (Basic scenarios covered)
- **Default Notification Template Sets**: ~65% (Basic + multi-scope)

### Target Test Coverage
- **Central Notification Channels**: 95%
- **Central Notification Rules**: 90%
- **Default Notification Template Sets**: 85%

## 🔄 CONTINUOUS TESTING STRATEGY

### Automated Testing
1. **Unit Tests**: All CRUD operations for each resource
2. **Integration Tests**: Cross-component functionality
3. **Regression Tests**: Version compatibility
4. **Performance Tests**: Load and stress testing

### Manual Testing
1. **UI Integration**: Terraform → Harness UI validation
2. **Real Notification Testing**: Actual email/Slack delivery
3. **Cross-Platform Testing**: Different Terraform versions

## 📝 TEST EXECUTION CHECKLIST

### Pre-Test Setup
- [ ] Test environment configuration
- [ ] Required Harness accounts and permissions
- [ ] Email/Slack/Teams test endpoints
- [ ] Terraform provider build

### Test Execution
- [ ] Run existing unit tests
- [ ] Execute new test additions
- [ ] Verify data source functionality
- [ ] Test import/export scenarios
- [ ] Validate error handling

### Post-Test Validation
- [ ] Resource cleanup verification
- [ ] No resource leaks
- [ ] Test report generation
- [ ] Coverage metrics update

## 🚀 IMPLEMENTATION TIMELINE

### Phase 1 (Week 1-2): High Priority Gaps
- Webhook and DataDog channel tests
- Custom template functionality
- User groups integration

### Phase 2 (Week 3-4): Medium Priority
- Multi-entity template tests
- Error handling scenarios
- Data source comprehensive testing

### Phase 3 (Week 5-6): Low Priority & Polish
- Performance and edge case tests
- Documentation updates
- CI/CD integration

This test plan provides a comprehensive roadmap for validating the Central Notification System implementation in the Harness Terraform Provider.