import sys

import psutil


def check_if_process_running(process_name):
    """
    Check if there is any running process that equals the given process_name.
    """
    # Iterate over the all the running process
    for proc in psutil.process_iter():
        try:
            # Check if process name is as the given name string.
            if process_name.lower() == proc.name().lower():
                return True
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass
    return False


if __name__ == '__main__':
    name = ''
    try:
        name = sys.argv[1]
    except KeyError:
        exit(1)
    if check_if_process_running(name):
        exit(0)
    exit(1)
