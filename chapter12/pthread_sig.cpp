#include <pthread.h>
#include <iostream>
#include <thread>
#include <signal.h>

int quit_flag;
sigset_t mask;
pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t wait = PTHREAD_COND_INITIALIZER;

void*
thr_fn(void *arg)
{
    int err, signo;

    for(;;)
    {
        err = sigwait(&mask, &signo);

        if (err != 0)
        {
            std::cout<<"sigwait error:" << err << std::endl;
            exit(-1);
        }
        
        switch (signo)
        {
        case SIGINT:
            std::cout<<"\ninterupt"<<std::endl;
            break;
        case SIGQUIT:
            pthread_mutex_lock(&lock);
            quit_flag = 1;
            pthread_mutex_unlock(&lock);
            pthread_cond_signal(&wait);
            return(0);
            break;
        default:
            break;
        }
    }
}


int main(int argc, char const *argv[])
{
    int err;
    sigset_t oldmask;

    pthread_t tid;

    sigemptyset(&mask);
    sigaddset(&mask, SIGINT);
    sigaddset(&mask, SIGQUIT);
    if (err = pthread_sigmask(SIG_BLOCK, &mask, &oldmask) != 0)
    {
        std::cout<<"SIG_BLOCK error"<<std::endl;
        exit(err);
    }

    
    err = pthread_create(&tid, nullptr, thr_fn, nullptr);
    if (err != 0 )
    {
        std::cout<<"creat thread error" << std::endl;
        exit(-1);
    }

    pthread_mutex_lock(&lock);
    while(!quit_flag)
    {
        pthread_cond_wait(&wait, &lock);
    }
    pthread_mutex_unlock(&lock);

    quit_flag = 0;

    if(sigprocmask(SIG_SETMASK, &oldmask, NULL) < 0)
    {
        std::cout<<"SIT_SETMASK error" << std::endl;
        exit(-1);
    }

    return 0;
}
