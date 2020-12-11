# casbin-rbac
trying out Casbin for RBAC

```
go run .
alice is a member of the following roles: [role:admin], and her permissions are: [[role:admin data read] [role:admin data write]]
bob is a member of the following roles: [role:user], and his permissions are: [[role:user data read]]
```

### Sources
- https://github.com/jeromefroe/go-casbin-example
