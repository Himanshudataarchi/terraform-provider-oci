// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_identity_domains_api_key", IdentityDomainsApiKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_api_keys", IdentityDomainsApiKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_auth_token", IdentityDomainsAuthTokenDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_auth_tokens", IdentityDomainsAuthTokensDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_authentication_factor_setting", IdentityDomainsAuthenticationFactorSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_authentication_factor_settings", IdentityDomainsAuthenticationFactorSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_customer_secret_key", IdentityDomainsCustomerSecretKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_customer_secret_keys", IdentityDomainsCustomerSecretKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_dynamic_resource_group", IdentityDomainsDynamicResourceGroupDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_dynamic_resource_groups", IdentityDomainsDynamicResourceGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_group", IdentityDomainsGroupDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_groups", IdentityDomainsGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_provider", IdentityDomainsIdentityProviderDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_providers", IdentityDomainsIdentityProvidersDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_kmsi_setting", IdentityDomainsKmsiSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_kmsi_settings", IdentityDomainsKmsiSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_api_key", IdentityDomainsMyApiKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_api_keys", IdentityDomainsMyApiKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_auth_token", IdentityDomainsMyAuthTokenDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_auth_tokens", IdentityDomainsMyAuthTokensDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_customer_secret_key", IdentityDomainsMyCustomerSecretKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_customer_secret_keys", IdentityDomainsMyCustomerSecretKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_device", IdentityDomainsMyDeviceDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_devices", IdentityDomainsMyDevicesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_groups", IdentityDomainsMyGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_oauth2client_credential", IdentityDomainsMyOAuth2ClientCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_oauth2client_credentials", IdentityDomainsMyOAuth2ClientCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_smtp_credential", IdentityDomainsMySmtpCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_smtp_credentials", IdentityDomainsMySmtpCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_support_account", IdentityDomainsMySupportAccountDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_support_accounts", IdentityDomainsMySupportAccountsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_trusted_user_agent", IdentityDomainsMyTrustedUserAgentDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_trusted_user_agents", IdentityDomainsMyTrustedUserAgentsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_user_db_credential", IdentityDomainsMyUserDbCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_user_db_credentials", IdentityDomainsMyUserDbCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth2client_credential", IdentityDomainsOAuth2ClientCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth2client_credentials", IdentityDomainsOAuth2ClientCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_password_policies", IdentityDomainsPasswordPoliciesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_password_policy", IdentityDomainsPasswordPolicyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_smtp_credential", IdentityDomainsSmtpCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_smtp_credentials", IdentityDomainsSmtpCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user", IdentityDomainsUserDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user_db_credential", IdentityDomainsUserDbCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user_db_credentials", IdentityDomainsUserDbCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_users", IdentityDomainsUsersDataSource())
}
