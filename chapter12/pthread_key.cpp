#include <thread>
#include <chrono>
#include <iostream>
#include <pthread.h>
#include <mutex>
using namespace std;


static pthread_key_t key;
static pthread_once_t flag = PTHREAD_ONCE_INIT;
pthread_mutex_t mm = PTHREAD_MUTEX_INITIALIZER;


void sleep(int seconds)
{
    this_thread::sleep_for(chrono::seconds(seconds));
}

void once()
{
    cout << "this is once func" << endl;
    pthread_key_create(&key, nullptr);
}

void show_info() 
{
    pthread_once(&flag, once);
    auto thread_id = this_thread::get_id();
    pthread_setspecific(key, static_cast<void*>(&thread_id));
    sleep(2);
    auto id = static_cast<thread::id*>(pthread_getspecific(key));
    pthread_mutex_lock(&mm);
    cout << "hello " << *id << endl;
    pthread_mutex_unlock(&mm);
}

int main(int argc, char const *argv[])
{
    thread t[10];
    for(int i = 0; i < 10; i++)
    {
        t[i] = thread(show_info);
    }
    for (auto &&tt : t)
    {
        tt.join();

    }

    pthread_key_delete(key);
    
    return 0;
}
