using Dapper;
using IdentityService.Data;
using IdentityService.Dtos.PermissionDtos;

namespace IdentityService.Repositories;
public class PermissionRepository
  (ISqlConnectionFactory connectionFactory)
 : IPermissionRepository
{
    public async Task<IEnumerable<PermissionDto>> GetPermissionsByRoleId(Guid id)
    {
        string query = """
           SELECT p.permission_id as PermissionId,
                  p.permission_name as PermissionName
           FROM permissions p
           INNER JOIN role_permissions rp ON p.permission_id = rp.permission_id
           WHERE rp.role_id = @Id
        """;
        using var connection = connectionFactory.Create();
        return await connection.QueryAsync<PermissionDto>(
            query,
            new { Id = id }
        );
    }
}


