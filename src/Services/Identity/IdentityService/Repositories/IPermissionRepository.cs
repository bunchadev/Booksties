using IdentityService.Dtos.PermissionDtos;

namespace IdentityService.Repositories;

public interface IPermissionRepository
{
    Task<IEnumerable<PermissionDto>> GetPermissionsByRoleId(Guid id);
}


