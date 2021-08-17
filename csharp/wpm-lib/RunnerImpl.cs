using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace wpm_lib
{
    class CommandRunnerImpl : CommandRunner
    {
        private RunnerContext mContext;

        public CommandRunnerImpl()
        {
            this.mContext = new RunnerContext();
        }

        void Runner.Go()
        {
            throw new NotImplementedException();
        }

        GuiRunner CommandRunner.GUI()
        {
            return new GuiRunnerImpl(this.mContext);
        }

        InstallerRunner CommandRunner.Install()
        {
            return new InstallerRunnerImpl(this.mContext);
        }

        ServiceRunner CommandRunner.StartService()
        {
            var runner = new ServiceRunnerImpl(this.mContext);
            return runner.DoStart();
        }

        ServiceRunner CommandRunner.StopService()
        {
            var runner = new ServiceRunnerImpl(this.mContext);
            return runner.DoStop();
        }
    }
}
