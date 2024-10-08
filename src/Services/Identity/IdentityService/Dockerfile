FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app
EXPOSE 8080

ENV ASPNETCORE_URLS=http://+:8080

USER app
FROM --platform=$BUILDPLATFORM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG configuration=Release
WORKDIR /src
COPY ["src/Services/Identity/IdentityService/IdentityService.csproj", "src/Services/Identity/IdentityService/"]
COPY ["src/CommonLib/CommonLib/CommonLib.csproj", "src/CommonLib/CommonLib/"]
RUN dotnet restore "src/Services/Identity/IdentityService/IdentityService.csproj"
COPY . .
WORKDIR "/src/src/Services/Identity/IdentityService"
RUN dotnet build "IdentityService.csproj" -c $configuration -o /app/build

FROM build AS publish
ARG configuration=Release
RUN dotnet publish "IdentityService.csproj" -c $configuration -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "IdentityService.dll"]
