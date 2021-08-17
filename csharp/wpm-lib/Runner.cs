using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace wpm_lib
{
    public interface Runner
    {

        void Go();
    }


    public interface CommandRunner : Runner
    {
        InstallerRunner Install();

        GuiRunner GUI();

        ServiceRunner StartService();

        ServiceRunner StopService();
    }


    public interface GuiRunner : Runner
    {
    }

    public interface InstallerRunner : Runner
    {
    }

    public interface ServiceRunner : Runner
    {
        ServiceRunner DoStop();
        ServiceRunner DoStart();
        ServiceRunner DoRestart();
    }



    public class WPM
    {
        public static CommandRunner Init()
        {
            return new CommandRunnerImpl();
        }
    }

}
