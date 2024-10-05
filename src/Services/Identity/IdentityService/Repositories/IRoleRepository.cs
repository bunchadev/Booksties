namespace IdentityService.Repositories
{
    public interface IRoleRepository
    {
        Task<Guid?> GetRoleIdWithName(string name);
        Task<Guid?> GetRoleWithId(Guid id);
    }
}
