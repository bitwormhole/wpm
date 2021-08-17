using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace wpm_lib
{
    class ServiceRunnerImpl : RunnerBase, ServiceRunner
    {
        public ServiceRunnerImpl(RunnerContext ctx) : base(ctx)
        {
        }


        public ServiceRunner DoRestart()
        {
            return this;
        }

        public ServiceRunner DoStart()
        {
            return this;
        }

        public ServiceRunner DoStop()
        {
            return this;
        }
    }
}
