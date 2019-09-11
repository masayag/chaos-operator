package chaosengine

import (
	"testing"

	litmuschaosv1alpha1 "github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestNewRunnerPodForCR(t *testing.T) {
	tests := map[string]struct {
		cr      *litmuschaosv1alpha1.ChaosEngine
		aUUID   types.UID
		aExList []string
		isErr   bool
	}{
		"Test Positive": {
			cr: &litmuschaosv1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-runner",
					Namespace: "test",
				},
			},
			aUUID:   "fake_id",
			aExList: []string{"exp-1"},
			isErr:   false,
		},
		"Test Negative-1": {
			cr: &litmuschaosv1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{},
			},
			aUUID:   "fake_id",
			aExList: []string{"exp-1"},
			isErr:   true,
		},
		"Test Negative-2 ": {
			cr: &litmuschaosv1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-runner",
					Namespace: "test",
				},
			},
			aUUID:   "",
			aExList: []string{"exp-1"},
			isErr:   true,
		},
		"Test Negative-3 ": {
			cr: &litmuschaosv1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-runner",
					Namespace: "test",
				},
			},
			aUUID:   "fake_id",
			aExList: []string{},
			isErr:   true,
		},
	}
	for name, mock := range tests {
		name, mock := name, mock
		t.Run(name, func(t *testing.T) {
			_, err := newRunnerPodForCR(mock.cr, mock.aUUID, mock.aExList)
			if mock.isErr && err == nil {
				t.Fatalf("Test %q failed: expected error not to be nil", name)
			}
			if !mock.isErr && err != nil {
				t.Fatalf("Test %q failed: expected error to be nil", name)
			}
		})
	}
}
func TestNewMonitorServiceForCR(t *testing.T) {
	tests := map[string]struct {
		cr    *litmuschaosv1alpha1.ChaosEngine
		isErr bool
	}{
		"Test Positive": {
			cr: &litmuschaosv1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-monitor",
					Namespace: "test",
				},
			},
			isErr: false,
		},
		"Test Negative": {
			cr: &litmuschaosv1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{},
			},
			isErr: true,
		},
	}
	for name, mock := range tests {
		name, mock := name, mock
		t.Run(name, func(t *testing.T) {

			_, err := newMonitorServiceForCR(mock.cr)
			if mock.isErr && err == nil {
				t.Fatalf("Test %q failed: expected error not to be nil", name)
			}
			if !mock.isErr && err != nil {
				t.Fatalf("Test %q failed: expected error to be nil", name)
			}
		})
	}
}