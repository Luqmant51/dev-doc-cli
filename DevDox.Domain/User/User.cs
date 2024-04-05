using DevDox.Domain.Abstraction;
using DevDox.Domain.User.Event;

namespace DevDox.Domain.User;

public sealed class User : Entity
{
    private User(Guid id, Username username) : base(id)
    {
        Username = username;
    }
    public Username Username { get; private set; }

    public static User Create(Username username)
    {
        var user = new User(Guid.NewGuid(), username);
        user.RaiseDomainEvent(new UserCreatedDomainEvent(user.Id));
        return user;
    }
}