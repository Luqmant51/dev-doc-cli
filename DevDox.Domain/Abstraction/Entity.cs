namespace DevDox.Domain.Abstraction;

public abstract class Entity
{
    private readonly List<IDomainEvent> _domainEvents = new();
    public Entity(Guid id)
    {
        Id = id;
    }
    public Guid Id { get; init; }

    public IReadOnlyList<IDomainEvent> GetDomainEvents()
    {
        return _domainEvents.ToList();
    }

    public void ClearDomainEvent()
    {
        _domainEvents.Clear();
    }

    protected void RaiseDomainEvent(IDomainEvent domainEvent)
    {
        _domainEvents.Add(domainEvent);
    }
}