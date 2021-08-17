using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Diagnostics;
using System.Linq;
using System.ServiceProcess;
using System.Text;
using wpm_lib;

namespace wpm_service
{
    public partial class WpmService : ServiceBase
    {
        public WpmService()
        {
            InitializeComponent();
        }

        protected override void OnStart(string[] args)
        {
            WPM.Init().StartService().Go();
        }

        protected override void OnStop()
        {
            WPM.Init().StopService().Go();
        }
    }
}
