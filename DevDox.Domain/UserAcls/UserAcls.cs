using DevDox.Domain.Abstraction;

namespace DevDox.Domain.UserAcls;

public class UserAcls : Entity
{
    public UserAcls(Guid id, User.User userId, WorkSpace.WorkSpace workSpaceId, Group.Group groupId) : base(id)
    {
        User_id = userId;
        WorkSpace_id = workSpaceId;
        Group_id = groupId;
    }
    public User.User User_id { get; private set; }
    public WorkSpace.WorkSpace WorkSpace_id { get; set; }
    public Group.Group Group_id { get; set; }

    public static UserAcls Create(User.User user_id, WorkSpace.WorkSpace workSpace, Group.Group group)
    {
        var UserAclsRecord = new UserAcls(Guid.NewGuid(), user_id, workSpace, group);
        return UserAclsRecord;
    }
}