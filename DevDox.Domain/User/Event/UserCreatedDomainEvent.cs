using DevDox.Domain.Abstraction;

namespace DevDox.Domain.User.Event;

public sealed record UserCreatedDomainEvent(Guid UserId) : IDomainEvent;