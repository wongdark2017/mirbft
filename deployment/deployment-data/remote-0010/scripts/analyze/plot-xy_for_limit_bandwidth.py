import matplotlib
matplotlib.use('Agg')

import matplotlib.pyplot as plt
import numpy as np

xpoints1 = np.array([1776.2674699564648,1787.970646418805,1342.6013875070346,1079.3972781344598])
ypoints1 = np.array([4157.9292011771,2769.496800214823,2136.9895155325444,557.7142221095335])

xpoints2 = np.array([2641.9546600257977,2646.296247979354,2239.48924323777,2154.824405765522])
ypoints2 = np.array([13020.584817170111,5476.111389805647,4756.964491448119,628.3828107630115])

xpoints3 = np.array([4132.354326453513,4239.216290144656,4732.005678180781])
ypoints3 = np.array([19316.55903239352,14059.777777111323,1598.1058919273944])

plt.plot(xpoints1, ypoints1,'-o')
plt.plot(xpoints2, ypoints2,'-o')
plt.plot(xpoints3, ypoints3,'-o')

xline1m = np.array([1776.2674699564648,2641.9546600257977])
yline1m = np.array([4157.9292011771,13020.584817170111])

xline1_5m = np.array([1787.970646418805,2646.296247979354,4132.354326453513])
yline1_5m = np.array([2769.496800214823,5476.111389805647,19316.55903239352])

xline2m = np.array([1342.6013875070346,2239.48924323777,4239.216290144656])
yline2m = np.array([2136.9895155325444,4756.964491448119,14059.777777111323])

xline10m = np.array([1079.3972781344598,2154.824405765522,4732.005678180781])
yline10m = np.array([557.7142221095335,628.3828107630115,1598.1058919273944])

plt.plot(xline1m, yline1m,'k--',linewidth=0.2)
plt.plot(xline1_5m, yline1_5m,'k--',linewidth=0.2)
plt.plot(xline2m, yline2m,'k--',linewidth=0.2)
plt.plot(xline10m, yline10m,'k--',linewidth=0.2)

plt.xlabel('throughput')
plt.ylabel('latency')
out_png = './out.png'
plt.savefig(out_png, dpi=400)
# plt.show()