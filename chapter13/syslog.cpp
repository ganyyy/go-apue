#include <syslog.h>
#include <stdarg.h>

int main(int argc, char const *argv[])
{
    openlog("MYLOG", LOG_PID|LOG_PERROR|LOG_CONS, LOG_SYSLOG);
    syslog(LOG_ERR, "open error for %s", __FILE__);
    return 0;
}
