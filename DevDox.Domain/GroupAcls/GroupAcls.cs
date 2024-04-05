using DevDox.Domain.Abstraction;

namespace DevDox.Domain.GroupAcls;

public sealed class GroupAcls : Entity
{
    public GroupAcls(Guid id,WorkSpace.WorkSpace workSpaceId, Group.Group groupId) : base(id)
    {
        WorkSpace_id = workSpaceId;
        Group_id = groupId;
    }
    public WorkSpace.WorkSpace WorkSpace_id { get; set; }
    public Group.Group Group_id { get; set; }

    public static GroupAcls Create(WorkSpace.WorkSpace workSpace, Group.Group group)
    {
        var GroupAclsRecord = new GroupAcls(Guid.NewGuid(), workSpace, group);
        return GroupAclsRecord;
    }
}