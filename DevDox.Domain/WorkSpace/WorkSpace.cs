using DevDox.Domain.Abstraction;
using DevDox.Domain.User.Event;
using DevDox.Domain.User;

namespace DevDox.Domain.WorkSpace;

public class WorkSpace : Entity
{
    public WorkSpace(Guid id, WorkSpaceName workSpaceName ) : base(id)
    {
        WorkSpaceName = workSpaceName;
    }

    public WorkSpaceName WorkSpaceName { get; private set; }

    public static WorkSpace Create(WorkSpaceName workSpaceName)
    {
        var workSpacename = new WorkSpace(Guid.NewGuid(), workSpaceName);
        return workSpacename;
    }
}