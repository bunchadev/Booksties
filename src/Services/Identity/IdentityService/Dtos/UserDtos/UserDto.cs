namespace IdentityService.Dtos.UserDtos
{
    public record UserDto
    (
        Guid UserId,
        string Email,
        string Password,
        string AuthMethod,
        bool IsActive,
        Guid RoleId,
        string RoleName
    );
}
