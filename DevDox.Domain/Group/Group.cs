using DevDox.Domain.Abstraction;

namespace DevDox.Domain.Group;

public class Group : Entity
{
    public Group(Guid id, GroupName groupName) : base(id)
    {
        GroupName = groupName;
    }
    public GroupName GroupName { get; private set; }
    public static Group Create(GroupName groupName)
    {
        var GroupType = new Group(Guid.NewGuid(), groupName);
        return GroupType;
    }
}