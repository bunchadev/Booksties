using IdentityService.Data;
using IdentityService.Services;
using Microsoft.AspNetCore.Authorization;

var builder = WebApplication.CreateBuilder(args);

var assembly = typeof(Program).Assembly;

builder.Services.AddMediatR(config =>
{
    config.RegisterServicesFromAssembly(assembly);
    config.AddOpenBehavior(typeof(CommandValidationBehavior<,>));
    config.AddOpenBehavior(typeof(QueryValidationBehavior<,>));
    config.AddOpenBehavior(typeof(LoggingBehavior<,>));
});

builder.Services.AddValidatorsFromAssembly(assembly);

builder.Services.AddCarter();

builder.Services.AddSingleton<ISqlConnectionFactory, SqlConnectionFactory>();

builder.Services.AddScoped<IUserRepository, UserRepository>();

builder.Services.AddScoped<IRoleRepository, RoleRepository>();

builder.Services.AddScoped<ITokenRepository, TokenRepository>();

builder.Services.AddScoped<IPermissionRepository, PermissionRepository>();

builder.Services.AddScoped<IPasswordService, PasswordService>();

builder.Services.AddScoped<IJwtService, JwtService>();

builder.Services.AddCustomJwtAuthentication(builder.Configuration);

builder.Services.AddAuthorization(AuthorizationPolicies.AddPolicies);

builder.Services.AddSingleton<IAuthorizationHandler, PermissionAuthorizationHandler>();

builder.Services.AddExceptionHandler<CustomExceptionHandler>();

var app = builder.Build();

app.MapCarter();

app.UseExceptionHandler(options => { });

app.UseAuthentication();
app.UseAuthorization();

app.Run();
public partial class Program { }


