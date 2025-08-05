**Title:** [Component/Feature] Short Description of the Change

**Summary:**
Provide a brief overview of what the PR does and why it is needed.

**Details:**
Explain the changes in detail, including any relevant context or background information.

**Related Issues:**
Reference any related issues or tickets 

**Testing Instructions:**
Describe how the changes were tested. Include any test cases or steps to verify the functionality.
Please include the below test scenario with your changes.
- Create the resource and execute terraform apply. (Verify the resource is created)
- Execute terraform apply again without any changes. (Verify no changes should be done)
- Update the resource and execute terraform apply. (Verify the resource updated successfully)
- Remove the content from resource file and execute terraform apply. (Verify the resource has been deleted)
- Add again the resource and execute the terraform apply. (Verify the resource is created)
- Verify the import the resource file is working fine.
- In case of remote entity, Verify for both default and non default branch.

**Screenshots:**
Include before and after screenshots to visually demonstrate the changes.
**Checklist:**
- [x] Code changes are well-documented.
- [x] Tests have been added/updated.
- [x] Changes have been tested locally.
- [ ] Documentation has been updated.

<details>
  <summary>PR Check triggers</summary>
  
- Build Test: `trigger build`
- Sub Category Field Check: `trigger subcategoryfieldcheck`
- Git Leaks: `trigger gitleaks`
- Message Check: `trigger messagecheck`
</details>