# OpenAPI spec has been modified

- Component schemas Limit, Offset and Curser received the 'Schema' prefix to prevent conflicts (i.e. LimitSchema)
- all "format": "date-time" has been removed, as go's default time library uses a different ISO
- remove anyOf, keep only support for project containers
