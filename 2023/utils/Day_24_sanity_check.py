import numpy as np
from numpy.linalg import inv


def crossMAtrix(a):
    return np.array([
        [0, -a[2], a[1]],
        [a[2], 0, -a[0]],
        [-a[1], a[0], 0]])


if __name__ == '__main__':
    h_0 = np.array([19, 13, 30])
    v_0 = np.array([-2, 1, -2])
    h_1 = np.array([18, 19, 22])
    v_1 = np.array([-1, -1, -2])
    h_2 = np.array([20, 25, 34])
    v_2 = np.array([-2, -2, -4])

    rhs = np.append(np.cross(-h_0, v_0) + np.cross(h_1, v_1),
                    np.cross(-h_1, v_1) + np.cross(h_2, v_2))
    print(rhs)

    m_0 = crossMAtrix(v_0) - crossMAtrix(v_1)
    m_1 = crossMAtrix(v_1) - crossMAtrix(v_2)

    m_2 = -crossMAtrix(h_0) + crossMAtrix(h_1)
    m_3 = -crossMAtrix(h_1) + crossMAtrix(h_2)

    m = np.concatenate((
        np.concatenate((m_0, m_1), axis=0),
        np.concatenate((m_2, m_3), axis=0)
    ), axis=1)
    print(m)
    m_inv = inv(m)
    print(np.dot(m_inv, rhs))
