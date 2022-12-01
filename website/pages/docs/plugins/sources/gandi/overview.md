# Gandi Plugin

The CloudQuery Gandi plugin pulls configuration out of Gandi resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Gandi, `cloudquery` needs to be authenticated. An API key is required for authentication. Get your API key from [Gandi's Account Settings Page](https://account.gandi.net/en/).

## Query Examples

### Get list of domains which will expire in 90 days

```sql
select fqdn, dates->>'registry_ends_at' as registry_ends_at, date_trunc('day', (dates->>'registry_ends_at')::timestamp - current_timestamp) as days_left from gandi_domains where ((dates->>'registry_ends_at')::timestamp - interval '90 day') < current_timestamp order by 1;
```

### Get domains in Gandi LiveDNS

```sql
select fqdn, current, nameservers from gandi_domain_live_dns order by 1;
```

### Get list of domain glue records

```sql
select fqdn, name, ips from gandi_domain_glue_records order by 1, 2;
```

### Get list of LiveDNS domains which don't have automatic snapshots enabled

```sql
select d.fqdn from gandi_livedns_domains d left join gandi_livedns_snapshots s on s.fqdn=d.fqdn and s.automatic where s.fqdn is null;
```
### Inspect LiveDNS snapshots for a given domain

```sql
select count(1) as number_of_snapshots, max(created_at) as last_snapshot_at from gandi_livedns_snapshots where fqdn = 'yourdomain.com';
```