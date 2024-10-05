namespace IdentityService.Dtos.PermissionDtos;

public record class PermissionDto
{
    public Guid PermissionId { get; set; }
    public string? PermissionName { get; set; }
}
