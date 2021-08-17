using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace wpm_lib
{

    class RunnerContext
    {
    }


    class RunnerBase : Runner
    {
        public RunnerBase(RunnerContext ctx)
        {
            this.Context = ctx;
        }

        public RunnerContext Context { get; }

        public void Go()
        {
           // throw new NotImplementedException();
        }
    }
}
