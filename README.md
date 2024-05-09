# json2ini

Takes a JSON document in the format

```json
{
    "section": {
        "key": "value"
    }
}
```

and transforms it into ini

```ini
[section]
key=value
```

This is perfect for generating rclone config files for example.
