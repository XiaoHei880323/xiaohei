# Windows Client

This directory contains a desktop launcher:

- `xiaohei-client.exe`: opens `http://127.0.0.1:8888/admin/page/index`.

## Behavior

1. The launcher checks `http://127.0.0.1:8888/admin/page/index` and `http://127.0.0.1:8888/from/xiaohe`.
2. If the backend is not running, it starts `api\api.exe -f etc\api-api.yaml`.
3. After the backend is ready, it opens `http://127.0.0.1:8888/admin/page/index`.
4. The page is rendered by Edge app mode, so the client shows the actual admin webpage content.
5. The desktop client stores its own login/session data in `windows-client\client-profile`.

## Requirements

- `windows-client` and `api` must remain sibling directories.
- `api\api.exe` and `api\etc\api-api.yaml` must exist.
- MySQL must be reachable, otherwise the backend cannot start successfully.
- Microsoft Edge is preferred; Google Chrome is used as fallback.

## Rebuild

```powershell
go build -ldflags "-H windowsgui" -o xiaohei-client.exe ./launcher
```
